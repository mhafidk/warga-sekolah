<script lang="ts">
  import { auth } from '$lib/auth.svelte';
  import { resolve } from '$app/paths';
</script>

<svelte:head>
  <title>Warga Sekolah — Portal Administrasi Sekolah Modern</title>
  <meta name="description" content="Portal administrasi yang minimalis, modern, dan aman untuk mengelola data warga sekolah, autentikasi, dan kredensial." />
</svelte:head>

<section class="max-w-7xl mx-auto px-6 py-20 md:py-32 grow flex flex-col justify-center animate-fade-in">
  <!-- Hero Header -->
  <div class="max-w-3xl mb-16">
    <div class="inline-flex items-center space-x-2 px-3 py-1 rounded-full border border-neutral-200 text-[11px] font-mono tracking-wider uppercase mb-6 bg-neutral-100/50">
      <span class="w-1.5 h-1.5 rounded-full bg-accent-blue animate-pulse"></span>
      <span>Status Sistem: Aktif</span>
    </div>

    <h1 class="font-display text-5xl md:text-7xl font-bold tracking-tight leading-[1.05] mb-6">
      Mengelola <br class="hidden sm:inline" />
      <span class="text-neutral-400 font-light">warga sekolah Anda,</span> <br />
      with total clarity.
    </h1>

    <p class="text-base md:text-lg text-neutral-500 leading-relaxed max-w-xl">
      Warga Sekolah adalah portal administrasi berkinerja tinggi. Kelola pengguna, verifikasi kredensial, dan amankan API internal sekolah melalui antarmuka minimalis.
    </p>

    <!-- CTAs -->
    <div class="mt-8 flex flex-wrap gap-4">
      {#if auth.initialized}
        {#if auth.token}
          <a
            href={resolve('/dashboard')}
            class="px-6 py-3 rounded-md text-sm font-semibold bg-neutral-900 text-white hover:bg-neutral-800 shadow-sm transition-all hover:-translate-y-px active:translate-y-0"
          >
            Ke Dasbor
          </a>
        {:else}
          <a
            href={resolve('/signup')}
            class="px-6 py-3 rounded-md text-sm font-semibold bg-neutral-900 text-white hover:bg-neutral-800 shadow-sm transition-all hover:-translate-y-px"
          >
            Mulai Sekarang
          </a>
          <a
            href={resolve('/login')}
            class="px-6 py-3 rounded-md text-sm font-semibold border border-neutral-300 hover:border-neutral-400 bg-transparent text-neutral-700 transition-all hover:-translate-y-px"
          >
            Masuk
          </a>
        {/if}
      {/if}
    </div>
  </div>

  <!-- Thin separating line -->
  <div class="h-px bg-neutral-200 w-full mb-16 relative overflow-hidden">
    <div class="absolute inset-0 bg-neutral-900 animate-draw-line" style="width: 100%"></div>
  </div>

  <!-- Features Bento Grid -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    <!-- Card 1 -->
    <div class="group p-8 rounded-lg border border-neutral-200 bg-neutral-50/50 hover:border-neutral-400 transition-all duration-300">
      <div class="font-mono text-xs opacity-50 mb-6">01 // AMAN</div>
      <h3 class="font-display text-lg font-semibold mb-2 group-hover:text-accent-blue transition-colors">
        Autentikasi JWT
      </h3>
      <p class="text-xs text-neutral-400 leading-relaxed">
        Verifikasi sesi cepat dan tanpa status yang dibangun dengan token standar industri HMAC-SHA256 dan kata sandi terenkripsi.
      </p>
    </div>

    <!-- Card 2 -->
    <div class="group p-8 rounded-lg border border-neutral-200 bg-neutral-50/50 hover:border-neutral-400 transition-all duration-300">
      <div class="font-mono text-xs opacity-50 mb-6">02 // CEPAT</div>
      <h3 class="font-display text-lg font-semibold mb-2 group-hover:text-accent-blue transition-colors">
        Arsitektur Tersemat
      </h3>
      <p class="text-xs text-neutral-400 leading-relaxed">
        Bundel frontend SvelteKit dikompilasi langsung menjadi aset statis, disajikan secara asli oleh backend Go melalui penyematan standar.
      </p>
    </div>

    <!-- Card 3 -->
    <div class="group p-8 rounded-lg border border-neutral-200 bg-neutral-50/50 hover:border-neutral-400 transition-all duration-300">
      <div class="font-mono text-xs opacity-50 mb-6">03 // BERSIH</div>
      <h3 class="font-display text-lg font-semibold mb-2 group-hover:text-accent-blue transition-colors">
        Minimalisme Murni
      </h3>
      <p class="text-xs text-neutral-400 leading-relaxed">
        Kegunaan tinggi, kesederhanaan visual, dan kepadatan informasi maksimum. Dirancang agar bersih dan mudah dibaca di semua layar.
      </p>
    </div>
  </div>
</section>
