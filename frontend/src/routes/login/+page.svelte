<script lang="ts">
  import { goto } from '$app/navigation';
  import { auth } from '$lib/auth.svelte';
  import { resolve } from '$app/paths';
  import { onMount } from 'svelte';

  let email = $state('');
  let password = $state('');
  let loading = $state(false);
  let errorMsg = $state('');

  onMount(() => {
    // If already logged in, send directly to dashboard
    if (auth.token) {
      goto(resolve('/dashboard'));
    }
  });

  async function handleLogin(e: SubmitEvent) {
    e.preventDefault();
    if (!email || !password) {
      errorMsg = 'Harap isi semua kolom.';
      return;
    }

    loading = true;
    errorMsg = '';

    try {
      const response = await fetch('/api/v1/users/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
      });

      if (!response.ok) {
        const text = await response.text();
        errorMsg = text || 'Email atau kata sandi salah.';
        loading = false;
        return;
      }

      const data = await response.json();
      if (data.token && data.user) {
        auth.login(data.token, data.user);
        goto(resolve('/dashboard'));
      } else {
        errorMsg = 'Menerima struktur respons yang tidak valid dari server.';
      }
    } catch (err) {
      errorMsg = 'Kesalahan jaringan. Pastikan server backend sedang berjalan.';
      console.error(err);
    } finally {
      loading = false;
    }
  }
</script>

<svelte:head>
  <title>Masuk — Warga Sekolah</title>
</svelte:head>

<div class="grow flex items-center justify-center px-6 py-20 animate-fade-in">
  <div class="w-full max-w-md">
    <!-- Header -->
    <div class="text-center mb-8">
      <div class="inline-flex items-center justify-center w-12 h-12 rounded border border-neutral-900 font-mono font-bold text-lg bg-neutral-900 text-white mb-4">
        WS
      </div>
      <h1 class="font-display text-2xl font-bold tracking-tight">Masuk ke portal administrasi</h1>
      <p class="text-xs text-neutral-400 mt-1">Masukkan kredensial Anda untuk mengelola basis data sekolah.</p>
    </div>

    <!-- Login Card -->
    <div class="p-8 rounded-lg border border-neutral-200 bg-neutral-50/50 backdrop-blur-sm">
      {#if errorMsg}
        <div class="mb-5 p-3.5 rounded border border-red-200 bg-red-50/55 text-red-600 text-xs font-mono flex items-start space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-alert-circle mt-0.5"><circle cx="12" cy="12" r="10"/><line x1="12" x2="12" y1="8" y2="12"/><line x1="12" x2="12" y1="16" y2="16"/></svg>
          <span class="flex-1">{errorMsg}</span>
        </div>
      {/if}

      <form onsubmit={handleLogin} class="space-y-4">
        <div>
          <label for="email" class="block text-[11px] font-mono tracking-wider uppercase text-neutral-400 mb-1.5">Alamat Email</label>
          <input
            type="email"
            id="email"
            bind:value={email}
            disabled={loading}
            placeholder="nama@sekolah.edu"
            required
            class="w-full px-3.5 py-2 text-sm rounded border border-neutral-200 bg-white text-neutral-900 placeholder-neutral-400 focus:border-neutral-900 transition-colors"
          />
        </div>

        <div>
          <div class="flex items-center justify-between mb-1.5">
            <label for="password" class="block text-[11px] font-mono tracking-wider uppercase text-neutral-400">Kata Sandi</label>
          </div>
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
            <span>Mengautentikasi...</span>
          {:else}
            <span>Masuk</span>
          {/if}
        </button>
      </form>
    </div>

    <!-- Redirect Link -->
    <p class="text-center text-xs text-neutral-400 mt-6">
      Belum punya akun administrator?
      <a href={resolve('/signup')} class="font-semibold text-neutral-900 hover:underline">Daftar di sini</a>.
    </p>
  </div>
</div>
