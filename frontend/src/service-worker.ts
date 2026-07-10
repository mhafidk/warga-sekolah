/// <reference lib="webworker" />

import { build, files, version } from '$service-worker';

const CACHE_NAME = `ws-cache-${version}`;
const ASSETS_TO_CACHE = [...build, ...files];

self.addEventListener('install', (event: any) => {
  event.waitUntil(
    caches
      .open(CACHE_NAME)
      .then((cache) => cache.addAll(ASSETS_TO_CACHE))
      .then(() => (self as any).skipWaiting())
  );
});

self.addEventListener('activate', (event: any) => {
  event.waitUntil(
    caches.keys().then(async (keys) => {
      // Delete old caches
      for (const key of keys) {
        if (key !== CACHE_NAME) {
          await caches.delete(key);
        }
      }
      await (self as any).clients.claim();
    })
  );
});

self.addEventListener('fetch', (event: any) => {
  const url = new URL(event.request.url);

  // Avoid intercepting non-GET requests or backend API routes
  if (event.request.method !== 'GET' || url.pathname.startsWith('/api/')) {
    return;
  }

  event.respondWith(
    caches.match(event.request).then((cachedResponse) => {
      if (cachedResponse) {
        return cachedResponse;
      }

      // If the asset is not cached and it's a page navigation request (e.g. /login or /dashboard)
      // serve the cached root index.html to allow SvelteKit client-side SPA routing to boot
      if (event.request.mode === 'navigate') {
        return caches.match('/index.html').then((indexResponse) => {
          if (indexResponse) return indexResponse;
          return fetch(event.request);
        });
      }

      return fetch(event.request).then((response) => {
        // Cache new static files if requested dynamically
        if (response.status === 200 && url.origin === location.origin) {
          const responseClone = response.clone();
          caches.open(CACHE_NAME).then((cache) => {
            cache.put(event.request, responseClone);
          });
        }
        return response;
      });
    })
  );
});
