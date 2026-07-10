<script lang="ts">
  import { goto } from '$app/navigation';
  import { auth } from '$lib/auth.svelte';
  import { resolve } from '$app/paths';
  import { onMount } from 'svelte';

  let fullName = $state('');
  let email = $state('');
  let password = $state('');
  let confirmPassword = $state('');
  let loading = $state(false);
  let errorMsg = $state('');
  let success = $state(false);
  let countdown = $state(3);

  onMount(() => {
    // If already logged in, redirect
    if (auth.token) {
      goto('/dashboard');
    }
  });

  function startRedirectCountdown() {
    const interval = setInterval(() => {
      countdown -= 1;
      if (countdown <= 0) {
        clearInterval(interval);
        goto('/login');
      }
    }, 1000);
  }

  async function handleRegister(e: SubmitEvent) {
    e.preventDefault();
    if (!fullName || !email || !password || !confirmPassword) {
      errorMsg = 'Semua kolom wajib diisi.';
      return;
    }

    if (password !== confirmPassword) {
      errorMsg = 'Kata sandi tidak cocok.';
      return;
    }

    if (password.length < 6) {
      errorMsg = 'Kata sandi minimal harus 6 karakter.';
      return;
    }

    loading = true;
    errorMsg = '';

    try {
      const response = await fetch('/api/v1/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          full_name: fullName,
          email,
          password
        })
      });

      if (!response.ok) {
        const text = await response.text();
        errorMsg = text || 'Pendaftaran gagal. Email mungkin sudah terdaftar.';
        loading = false;
        return;
      }

      success = true;
      startRedirectCountdown();
    } catch (err) {
      errorMsg = 'Kesalahan jaringan. Pastikan server backend sedang berjalan.';
      console.error(err);
    } finally {
      loading = false;
    }
  }
</script>

<svelte:head>
  <title>Daftar Akun — Warga Sekolah</title>
</svelte:head>

<div class="flex-grow flex items-center justify-center px-6 py-20 animate-fade-in">
  <div class="w-full max-w-md">
    <!-- Header -->
    <div class="text-center mb-8">
      <div class="inline-flex items-center justify-center w-12 h-12 rounded border border-neutral-900 font-mono font-bold text-lg bg-neutral-900 text-white mb-4">
        WS
      </div>
      <h1 class="font-display text-2xl font-bold tracking-tight">Buat akun administrator</h1>
      <p class="text-xs text-neutral-400 mt-1">Daftar sebagai administrator sistem untuk mengelola warga sekolah.</p>
    </div>

    <!-- Signup Card / Success Block -->
    <div class="p-8 rounded-lg border border-neutral-200 bg-neutral-50/50 backdrop-blur-sm">
      {#if success}
        <div class="text-center py-6 space-y-4">
          <div class="inline-flex items-center justify-center w-12 h-12 rounded-full border border-green-200 bg-green-55 text-green-600">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-check"><path d="M20 6 9 17l-5-5"/></svg>
          </div>
          <h3 class="font-display text-lg font-bold">Pendaftaran Berhasil</h3>
          <p class="text-xs text-neutral-400 leading-relaxed">
            Profil admin Anda telah terdaftar. Anda akan dialihkan ke halaman masuk dalam <span class="font-mono font-bold text-neutral-900">{countdown}</span> detik...
          </p>
          <div class="pt-4">
            <a href={resolve('/login')} class="text-xs font-semibold underline text-neutral-900">Alihkan sekarang</a>
          </div>
        </div>
      {:else}
        {#if errorMsg}
          <div class="mb-5 p-3.5 rounded border border-red-200 bg-red-50/55 text-red-600 text-xs font-mono flex items-start space-x-2">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-alert-circle mt-0.5"><circle cx="12" cy="12" r="10"/><line x1="12" x2="12" y1="8" y2="12"/><line x1="12" x2="12" y1="16" y2="16"/></svg>
            <span class="flex-1">{errorMsg}</span>
          </div>
        {/if}

        <form onsubmit={handleRegister} class="space-y-4">
          <div>
            <label for="fullName" class="block text-[11px] font-mono tracking-wider uppercase text-neutral-400 mb-1.5">Nama Lengkap</label>
            <input 
              type="text" 
              id="fullName" 
              bind:value={fullName}
              disabled={loading}
              placeholder="Nama Administrator" 
              required
              class="w-full px-3.5 py-2 text-sm rounded border border-neutral-200 bg-white text-neutral-900 placeholder-neutral-400 focus:border-neutral-900 transition-colors"
            />
          </div>

          <div>
            <label for="email" class="block text-[11px] font-mono tracking-wider uppercase text-neutral-400 mb-1.5">Alamat Email</label>
            <input 
              type="email" 
              id="email" 
              bind:value={email}
              disabled={loading}
              placeholder="admin@sekolah.edu" 
              required
              class="w-full px-3.5 py-2 text-sm rounded border border-neutral-200 bg-white text-neutral-900 placeholder-neutral-400 focus:border-neutral-900 transition-colors"
            />
          </div>

          <div>
            <label for="password" class="block text-[11px] font-mono tracking-wider uppercase text-neutral-400 mb-1.5">Kata Sandi</label>
            <input 
              type="password" 
              id="password" 
              bind:value={password}
              disabled={loading}
              placeholder="••••••••" 
              required
              class="w-full px-3.5 py-2 text-sm rounded border border-neutral-200 bg-white text-neutral-900 placeholder-neutral-400 focus:border-neutral-900 transition-colors"
            />
          </div>

          <div>
            <label for="confirmPassword" class="block text-[11px] font-mono tracking-wider uppercase text-neutral-400 mb-1.5">Konfirmasi Kata Sandi</label>
            <input 
              type="password" 
              id="confirmPassword" 
              bind:value={confirmPassword}
              disabled={loading}
              placeholder="••••••••" 
              required
              class="w-full px-3.5 py-2 text-sm rounded border border-neutral-200 bg-white text-neutral-900 placeholder-neutral-400 focus:border-neutral-900 transition-colors"
            />
          </div>

          <button 
            type="submit" 
            disabled={loading}
            class="w-full py-2.5 rounded text-sm font-semibold bg-neutral-900 text-white hover:bg-neutral-800 transition-colors shadow-sm flex items-center justify-center space-x-2 disabled:opacity-50"
          >
            {#if loading}
              <svg class="animate-spin h-4 w-4 text-current" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <span>Membuat Akun...</span>
            {:else}
              <span>Daftar</span>
            {/if}
          </button>
        </form>
      {/if}
    </div>

    <!-- Redirect Link -->
    {#if !success}
      <p class="text-center text-xs text-neutral-400 mt-6">
        Sudah punya akun administrator? 
        <a href={resolve('/login')} class="font-semibold text-neutral-900 hover:underline">Masuk di sini</a>.
      </p>
    {/if}
  </div>
</div>
