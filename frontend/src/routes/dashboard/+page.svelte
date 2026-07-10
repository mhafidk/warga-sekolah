<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { auth } from '$lib/auth.svelte';

  // State Management (Svelte 5 Runes)
  let loadingProfile = $state(true);
  let showAddModal = $state(false);
  let searchQuery = $state('');
  let activeTab = $state<'All' | 'Siswa' | 'Guru' | 'Staf'>('All');

  // Input states for new citizen
  let newName = $state('');
  let newEmail = $state('');
  let newRole = $state<'Siswa' | 'Guru' | 'Staf'>('Siswa');
  let newClassId = $state('10-A');
  let formError = $state('');

  // Local/Mock Database of School Citizens (warga) in Indonesian
  let citizensList = $state([
    { id: '1001', name: 'Ahmad Rafli', email: 'ahmad.rafli@school.edu', role: 'Siswa', classId: '12-IPA-1', status: 'Aktif' },
    { id: '1002', name: 'Sarah Callista', email: 'sarah.c@school.edu', role: 'Siswa', classId: '11-IPS-2', status: 'Aktif' },
    { id: '2001', name: 'Dwi Cahyono, M.Pd.', email: 'dwi.cahyono@school.edu', role: 'Guru', classId: 'Dep. Fisika', status: 'Aktif' },
    { id: '2002', name: 'Rina Herawati, S.S.', email: 'rina.hera@school.edu', role: 'Guru', classId: 'Dep. Inggris', status: 'Aktif' },
    { id: '3001', name: 'Budi Santoso', email: 'budi.s@school.edu', role: 'Staf', classId: 'Helpdesk TI', status: 'Aktif' },
    { id: '1003', name: 'Farhan Maulana', email: 'farhan.m@school.edu', role: 'Siswa', classId: '10-C', status: 'Ditangguhkan' },
    { id: '1004', name: 'Nadia Az-Zahra', email: 'nadia.az@school.edu', role: 'Siswa', classId: '12-IPA-3', status: 'Aktif' },
    { id: '2003', name: 'Heri Prasetyo, M.T.', email: 'heri.p@school.edu', role: 'Guru', classId: 'Dep. Matematika', status: 'Aktif' }
  ]);

  onMount(async () => {
    if (!auth.token) {
      goto('/login');
      return;
    }

    // Call user endpoint to verify token and load identity
    const success = await auth.checkMe();
    if (!success) {
      goto('/login');
      return;
    }
    loadingProfile = false;
  });

  // Derived filtered items based on search and selected filter tab
  const filteredCitizens = $derived(
    citizensList.filter(c => {
      const matchesSearch = c.name.toLowerCase().includes(searchQuery.toLowerCase()) || 
                            c.email.toLowerCase().includes(searchQuery.toLowerCase()) ||
                            c.id.includes(searchQuery) ||
                            c.classId.toLowerCase().includes(searchQuery.toLowerCase());
      const matchesTab = activeTab === 'All' ? true : c.role === activeTab;
      return matchesSearch && matchesTab;
    })
  );

  // Stats Derived
  const totalCount = $derived(citizensList.length);
  const studentCount = $derived(citizensList.filter(c => c.role === 'Siswa').length);
  const teacherCount = $derived(citizensList.filter(c => c.role === 'Guru').length);
  const staffCount = $derived(citizensList.filter(c => c.role === 'Staf').length);

  function addNewCitizen(e: SubmitEvent) {
    e.preventDefault();
    if (!newName || !newEmail) {
      formError = 'Harap isi semua kolom yang diperlukan.';
      return;
    }

    // Create custom unique ID
    const baseId = newRole === 'Siswa' ? 1000 : newRole === 'Guru' ? 2000 : 3000;
    const existingRoleList = citizensList.filter(c => c.role === newRole);
    const newId = (baseId + existingRoleList.length + 1).toString();

    const newMember = {
      id: newId,
      name: newName,
      email: newEmail,
      role: newRole,
      classId: newRole === 'Siswa' ? newClassId : newRole === 'Guru' ? `${newClassId}` : newClassId,
      status: 'Aktif'
    };

    citizensList = [newMember, ...citizensList];
    
    // Clear inputs and close modal
    newName = '';
    newEmail = '';
    newRole = 'Siswa';
    newClassId = '10-A';
    formError = '';
    showAddModal = false;
  }

  function removeCitizen(id: string) {
    citizensList = citizensList.filter(c => c.id !== id);
  }
</script>

<svelte:head>
  <title>Dasbor Administrasi — Warga Sekolah</title>
</svelte:head>

<div class="max-w-7xl mx-auto px-6 py-10 flex-grow w-full flex flex-col animate-fade-in">
  {#if loadingProfile}
    <div class="flex-grow flex flex-col items-center justify-center space-y-4 py-32">
      <svg class="animate-spin h-8 w-8 text-neutral-900" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="3"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <span class="text-xs font-mono tracking-widest uppercase opacity-60">Mengambil Profil Aman...</span>
    </div>
  {:else}
    <!-- Dashboard Header -->
    <div class="flex flex-col md:flex-row md:items-end justify-between border-b border-neutral-200 pb-8 mb-10">
      <div>
        <div class="font-mono text-xs opacity-50 mb-2">// SESI AKSES AMAN</div>
        <h1 class="font-display text-3xl md:text-4xl font-bold tracking-tight">
          Selamat datang kembali, {auth.user?.full_name || 'Admin'}
        </h1>
        <p class="text-xs text-neutral-400 mt-1.5">
          Peran: Administrator Sistem &bull; Server Sesi: Node Lokal
        </p>
      </div>

      <div class="mt-4 md:mt-0">
        <button 
          onclick={() => showAddModal = true}
          class="px-4 py-2 rounded text-xs font-semibold bg-neutral-900 text-white hover:bg-neutral-800 transition-colors shadow-sm inline-flex items-center space-x-1.5"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-plus"><path d="M5 12h14"/><path d="M12 5v19"/></svg>
          <span>Daftarkan Warga Sekolah</span>
        </button>
      </div>
    </div>

    <!-- Quick Stats Cards (Neo-Swiss Grid) -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-6 mb-10">
      <!-- Stat 1 -->
      <div class="p-6 rounded-lg border border-neutral-200 bg-neutral-50/50 backdrop-blur-sm">
        <div class="text-[10px] font-mono tracking-wider uppercase opacity-50 mb-4">Total Direktori</div>
        <div class="flex items-baseline space-x-2">
          <span class="font-display text-3xl font-bold font-mono">{totalCount}</span>
          <span class="text-xs text-neutral-400">warga</span>
        </div>
      </div>
      <!-- Stat 2 -->
      <div class="p-6 rounded-lg border border-neutral-200 bg-neutral-50/50 backdrop-blur-sm">
        <div class="text-[10px] font-mono tracking-wider uppercase opacity-50 mb-4">Siswa Terdaftar</div>
        <div class="flex items-baseline space-x-2">
          <span class="font-display text-3xl font-bold font-mono">{studentCount}</span>
          <span class="text-xs text-neutral-400">aktif</span>
        </div>
      </div>
      <!-- Stat 3 -->
      <div class="p-6 rounded-lg border border-neutral-200 bg-neutral-50/50 backdrop-blur-sm">
        <div class="text-[10px] font-mono tracking-wider uppercase opacity-50 mb-4">Guru Aktif</div>
        <div class="flex items-baseline space-x-2">
          <span class="font-display text-3xl font-bold font-mono">{teacherCount}</span>
          <span class="text-xs text-neutral-400">pengajar</span>
        </div>
      </div>
      <!-- Stat 4 -->
      <div class="p-6 rounded-lg border border-neutral-200 bg-neutral-50/50 backdrop-blur-sm">
        <div class="text-[10px] font-mono tracking-wider uppercase opacity-50 mb-4">Staf Operasional</div>
        <div class="flex items-baseline space-x-2">
          <span class="font-display text-3xl font-bold font-mono">{staffCount}</span>
          <span class="text-xs text-neutral-400">pendukung</span>
        </div>
      </div>
    </div>

    <!-- Interactive Table Section -->
    <div class="flex-grow flex flex-col">
      <!-- Filters and Search Bar -->
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-6">
        <!-- Tab selector -->
        <div class="flex border-b border-neutral-200 text-xs font-semibold">
          <button 
            onclick={() => activeTab = 'All'}
            class="px-4 py-2.5 transition-colors border-b-2 -mb-[2px] {activeTab === 'All' ? 'border-neutral-900 text-neutral-900' : 'border-transparent text-neutral-400 hover:text-neutral-600'}"
          >
            Semua
          </button>
          <button 
            onclick={() => activeTab = 'Siswa'}
            class="px-4 py-2.5 transition-colors border-b-2 -mb-[2px] {activeTab === 'Siswa' ? 'border-neutral-900 text-neutral-900' : 'border-transparent text-neutral-400 hover:text-neutral-600'}"
          >
            Siswa
          </button>
          <button 
            onclick={() => activeTab = 'Guru'}
            class="px-4 py-2.5 transition-colors border-b-2 -mb-[2px] {activeTab === 'Guru' ? 'border-neutral-900 text-neutral-900' : 'border-transparent text-neutral-400 hover:text-neutral-600'}"
          >
            Guru
          </button>
          <button 
            onclick={() => activeTab = 'Staf'}
            class="px-4 py-2.5 transition-colors border-b-2 -mb-[2px] {activeTab === 'Staf' ? 'border-neutral-900 text-neutral-900' : 'border-transparent text-neutral-400 hover:text-neutral-600'}"
          >
            Staf
          </button>
        </div>

        <!-- Search Input -->
        <div class="relative w-full md:max-w-xs">
          <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none opacity-40">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-search"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
          </span>
          <input 
            type="text" 
            bind:value={searchQuery}
            placeholder="Cari nama, ID, atau kelas..."
            class="w-full pl-9 pr-4 py-2 text-xs rounded border border-neutral-200 bg-neutral-50 text-neutral-900 placeholder-neutral-400 focus:border-neutral-900 transition-colors"
          />
        </div>
      </div>

      <!-- Directory Table -->
      <div class="w-full overflow-x-auto rounded border border-neutral-200 bg-neutral-50/20 backdrop-blur-sm">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="border-b border-neutral-200 bg-neutral-100/50 text-[10px] font-mono tracking-wider uppercase text-neutral-400">
              <th class="py-3.5 px-6">Kode ID</th>
              <th class="py-3.5 px-6">Nama Lengkap</th>
              <th class="py-3.5 px-6">Alamat Email</th>
              <th class="py-3.5 px-6">Peran</th>
              <th class="py-3.5 px-6">Kelas / Tugas</th>
              <th class="py-3.5 px-6">Status</th>
              <th class="py-3.5 px-6 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-neutral-200 text-xs">
            {#each filteredCitizens as citizen (citizen.id)}
              <tr class="hover:bg-neutral-100/40 transition-colors">
                <td class="py-4 px-6 font-mono text-neutral-400">#{citizen.id}</td>
                <td class="py-4 px-6 font-semibold">{citizen.name}</td>
                <td class="py-4 px-6 text-neutral-500">{citizen.email}</td>
                <td class="py-4 px-6">
                  <span class="px-2 py-0.5 rounded text-[10px] font-mono border {
                    citizen.role === 'Siswa' ? 'border-blue-200 bg-blue-50/30 text-blue-600' :
                    citizen.role === 'Guru' ? 'border-purple-200 bg-purple-50/30 text-purple-600' :
                    'border-neutral-300 bg-neutral-100/30 text-neutral-600'
                  }">
                    {citizen.role}
                  </span>
                </td>
                <td class="py-4 px-6 font-mono text-[11px]">{citizen.classId}</td>
                <td class="py-4 px-6">
                  <span class="inline-flex items-center space-x-1.5">
                    <span class="w-1.5 h-1.5 rounded-full {citizen.status === 'Aktif' ? 'bg-green-500' : 'bg-amber-500'}"></span>
                    <span class="text-[10px] opacity-75">{citizen.status}</span>
                  </span>
                </td>
                <td class="py-4 px-6 text-right">
                  <button 
                    onclick={() => removeCitizen(citizen.id)}
                    class="text-[10px] font-mono text-neutral-400 hover:text-red-500 transition-colors"
                  >
                    [HAPUS]
                  </button>
                </td>
              </tr>
            {:else}
              <tr>
                <td colspan="7" class="py-12 text-center text-xs text-neutral-400 font-mono">
                  TIDAK ADA DATA DIREKTORI YANG COCOK DENGAN PENCARIAN //
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  {/if}
</div>

<!-- Modal Form overlay (Stark Bauhaus panel) -->
{#if showAddModal}
  <div class="fixed inset-0 z-50 flex items-center justify-center p-6">
    <!-- Backdrop -->
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div class="absolute inset-0 bg-neutral-950/40 backdrop-blur-sm" onclick={() => showAddModal = false}></div>
    
    <!-- Panel -->
    <div class="relative w-full max-w-md p-8 rounded-lg border border-neutral-200 bg-neutral-50 text-neutral-900 shadow-xl animate-fade-in">
      <div class="flex items-center justify-between border-b border-neutral-200 pb-4 mb-6">
        <div>
          <h3 class="font-display text-lg font-bold">Daftarkan Warga Sekolah</h3>
          <p class="text-[10px] text-neutral-400 mt-0.5">Tambahkan catatan baru ke basis data direktori.</p>
        </div>
        <button 
          onclick={() => showAddModal = false}
          aria-label="Tutup Modal"
          class="text-neutral-400 hover:text-neutral-600 transition-colors"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-x"><line x1="18" x2="6" y1="6" y2="18"/><line x1="6" x2="18" y1="6" y2="18"/></svg>
        </button>
      </div>

      {#if formError}
        <div class="mb-4 p-2.5 rounded border border-red-200 bg-red-50 text-red-600 text-xs font-mono">
          {formError}
        </div>
      {/if}

      <form onsubmit={addNewCitizen} class="space-y-4">
        <div>
          <label for="newName" class="block text-[10px] font-mono tracking-wider uppercase text-neutral-400 mb-1">Nama Lengkap Warga</label>
          <input 
            type="text" 
            id="newName"
            bind:value={newName}
            placeholder="Contoh: Ahmad Rafli" 
            required
            class="w-full px-3 py-2 text-xs rounded border border-neutral-200 bg-white text-neutral-950 focus:border-neutral-900 transition-colors"
          />
        </div>

        <div>
          <label for="newEmail" class="block text-[10px] font-mono tracking-wider uppercase text-neutral-400 mb-1">Alamat Email</label>
          <input 
            type="email" 
            id="newEmail"
            bind:value={newEmail}
            placeholder="contoh@sekolah.edu" 
            required
            class="w-full px-3 py-2 text-xs rounded border border-neutral-200 bg-white text-neutral-950 focus:border-neutral-900 transition-colors"
          />
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label for="newRole" class="block text-[10px] font-mono tracking-wider uppercase text-neutral-400 mb-1">Peran Warga</label>
            <select 
              id="newRole"
              bind:value={newRole}
              class="w-full px-3 py-2 text-xs rounded border border-neutral-200 bg-white text-neutral-950 focus:border-neutral-900 transition-colors"
            >
              <option value="Siswa">Siswa</option>
              <option value="Guru">Guru</option>
              <option value="Staf">Staf</option>
            </select>
          </div>

          <div>
            <label for="newClass" class="block text-[10px] font-mono tracking-wider uppercase text-neutral-400 mb-1">
              {#if newRole === 'Siswa'}Penugasan Kelas{:else if newRole === 'Guru'}Departemen{:else}Departemen / Tugas{/if}
            </label>
            {#if newRole === 'Siswa'}
              <select 
                id="newClass"
                bind:value={newClassId}
                class="w-full px-3 py-2 text-xs rounded border border-neutral-200 bg-white text-neutral-950 focus:border-neutral-900 transition-colors"
              >
                <option value="10-A">10-A</option>
                <option value="10-B">10-B</option>
                <option value="11-IPA-1">11-IPA-1</option>
                <option value="11-IPS-2">11-IPS-2</option>
                <option value="12-IPA-1">12-IPA-1</option>
                <option value="12-IPS-1">12-IPS-1</option>
              </select>
            {:else}
              <input 
                type="text" 
                id="newClass"
                bind:value={newClassId}
                placeholder={newRole === 'Guru' ? 'Contoh: Matematika' : 'Contoh: Administrasi'}
                required
                class="w-full px-3 py-2 text-xs rounded border border-neutral-200 bg-white text-neutral-950 focus:border-neutral-900 transition-colors"
              />
            {/if}
          </div>
        </div>

        <div class="pt-4 flex justify-end space-x-3">
          <button 
            type="button" 
            onclick={() => showAddModal = false}
            class="px-4 py-2 rounded text-xs font-semibold border border-neutral-300 bg-transparent text-neutral-700 hover:bg-neutral-100 transition-colors"
          >
            Batal
          </button>
          <button 
            type="submit" 
            class="px-4 py-2 rounded text-xs font-semibold bg-neutral-900 text-white hover:bg-neutral-800 transition-colors"
          >
            Tambah Catatan
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}
