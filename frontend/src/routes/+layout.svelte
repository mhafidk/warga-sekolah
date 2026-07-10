<script lang="ts">
  import './layout.css';
  import { onMount } from 'svelte';
  import { page } from '$app/state';
  import { auth } from '$lib/auth.svelte';
  import { resolve } from '$app/paths';
  import favicon from '$lib/assets/favicon.svg';

  let { children } = $props();

  onMount(() => {
    // Check user auth token status
    if (auth.token) {
      auth.checkMe();
    }

    // Register PWA Service Worker
    if ('serviceWorker' in navigator) {
      navigator.serviceWorker.register('/service-worker.js', { type: 'module' })
        .then(() => console.log('Service Worker registered successfully'))
        .catch(err => console.warn('Service Worker registration failed:', err));
    }
  });

  // Derived variable for active route styling
  const currentPath = $derived(page.url.pathname);
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
  <title>Warga Sekolah — Sistem Manajemen Sekolah</title>
</svelte:head>

<div class="min-h-screen flex flex-col font-sans bg-neutral-50 text-neutral-900 selection:bg-neutral-900 selection:text-white">
  <!-- Top Navigation -->
  <header class="sticky top-0 z-40 w-full border-b border-neutral-200 bg-neutral-50/80 backdrop-blur-md">
    <div class="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
      <!-- Logo -->
      <a href={resolve('/')} class="flex items-center space-x-3 group">
        <div class="w-8 h-8 rounded border border-neutral-900 flex items-center justify-center font-mono font-bold text-sm bg-neutral-900 text-white transition-all group-hover:-rotate-6">
          WS
        </div>
        <span class="font-display font-medium text-lg tracking-tight hover:opacity-85 transition-opacity">
          warga<span class="text-neutral-400">_sekolah</span>
        </span>
      </a>

      <!-- Nav Links -->
      <nav class="hidden md:flex items-center space-x-8 text-sm font-medium">
        <a href={resolve('/')} class="transition-colors hover:text-neutral-500 {currentPath === '/' ? 'text-neutral-900 border-b border-neutral-900 pb-0.5' : 'text-neutral-500'}">
          Beranda
        </a>
        {#if auth.token}
          <a href={resolve('/dashboard')} class="transition-colors hover:text-neutral-500 {currentPath.startsWith('/dashboard') ? 'text-neutral-900 border-b border-neutral-900 pb-0.5' : 'text-neutral-500'}">
            Dasbor
          </a>
        {/if}
      </nav>

      <!-- Right actions -->
      <div class="flex items-center space-x-4">
        <!-- Auth Actions -->
        {#if auth.initialized}
          {#if auth.token && auth.user}
            <div class="hidden sm:flex items-center space-x-3 text-sm">
              <span class="font-mono text-xs opacity-60">@{auth.user.email.split('@')[0]}</span>
              <button
                onclick={() => auth.logout()}
                class="px-3.5 py-1.5 rounded-md text-xs font-semibold bg-neutral-900 text-white hover:bg-neutral-800 transition-colors"
              >
                Keluar
              </button>
            </div>
          {:else}
            <div class="flex items-center space-x-3">
              <a
                href={resolve('/login')}
                class="px-3 py-1.5 text-xs font-semibold text-neutral-600 hover:text-neutral-950 transition-colors"
              >
                Masuk
              </a>
              <a
                href={resolve('/signup')}
                class="px-3.5 py-1.5 rounded-md text-xs font-semibold bg-neutral-900 text-white hover:bg-neutral-800 transition-colors"
              >
                Daftar
              </a>
            </div>
          {/if}
        {/if}
      </div>
    </div>
  </header>

  <!-- Mobile Nav Subheader (Only visible on small screens when logged in) -->
  {#if auth.token && auth.initialized}
    <div class="md:hidden border-b border-neutral-200 bg-neutral-50/50 backdrop-blur">
      <div class="flex items-center justify-around h-10 text-xs font-semibold">
        <a href={resolve('/')} class="text-neutral-500 hover:text-neutral-900 {currentPath === '/' ? 'text-neutral-900 underline' : ''}">
          Beranda
        </a>
        <a href={resolve('/dashboard')} class="text-neutral-500 hover:text-neutral-900 {currentPath.startsWith('/dashboard') ? 'text-neutral-900 underline' : ''}">
          Dasbor
        </a>
        <button onclick={() => auth.logout()} class="text-neutral-500 hover:text-red-500 transition-colors">
          Keluar
        </button>
      </div>
    </div>
  {/if}

  <!-- Main Viewport -->
  <main class="grow flex flex-col">
    {@render children()}
  </main>

  <!-- Footer -->
  <footer class="border-t border-neutral-200 py-12 bg-neutral-50/50">
    <div class="max-w-7xl mx-auto px-6 grid grid-cols-1 md:grid-cols-3 gap-8">
      <div>
        <div class="flex items-center space-x-2 mb-3">
          <div class="w-5 h-5 rounded border border-neutral-900 flex items-center justify-center font-mono font-bold text-[9px] bg-neutral-900 text-white">
            WS
          </div>
          <span class="font-display font-medium text-sm tracking-tight">warga_sekolah</span>
        </div>
        <p class="text-xs text-neutral-400 leading-relaxed max-w-sm">
          Gerbang administrasi sekolah yang minimalis. Kelola siswa, guru, staf, dan pengguna akademik melalui aplikasi modern dan aman.
        </p>
      </div>
      <div class="flex flex-col space-y-2 md:items-end md:col-span-2 text-xs">
        <div class="font-mono text-neutral-400">
          DIRANCANG DENGAN SENI TINGGI // MINIMALISME NEO-SWISS
        </div>
        <div class="text-neutral-400">
          © {new Date().getFullYear()} Warga Sekolah. Hak Cipta Dilindungi.
        </div>
      </div>
    </div>
  </footer>
</div>
