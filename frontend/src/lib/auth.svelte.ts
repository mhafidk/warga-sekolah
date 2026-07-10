import { goto } from '$app/navigation';

export class AuthState {
  token = $state<string | null>(null);
  user = $state<any | null>(null);
  initialized = $state<boolean>(false);

  constructor() {
    if (typeof window !== 'undefined') {
      this.token = localStorage.getItem('ws_token');
      const savedUser = localStorage.getItem('ws_user');
      if (savedUser) {
        try {
          this.user = JSON.parse(savedUser);
        } catch {
          this.user = null;
        }
      }
      this.initialized = true;
    }
  }

  login(token: string, user: any) {
    this.token = token;
    this.user = user;
    localStorage.setItem('ws_token', token);
    localStorage.setItem('ws_user', JSON.stringify(user));
  }

  logout() {
    this.token = null;
    this.user = null;
    localStorage.removeItem('ws_token');
    localStorage.removeItem('ws_user');
    goto('/login');
  }

  async checkMe() {
    if (!this.token) return false;
    try {
      const res = await fetch('/api/v1/users/me', {
        headers: {
          'Authorization': `Bearer ${this.token}`
        }
      });
      if (res.ok) {
        const userData = await res.json();
        this.user = userData;
        localStorage.setItem('ws_user', JSON.stringify(userData));
        return true;
      } else {
        this.logout();
        return false;
      }
    } catch (e) {
      console.error('Failed to authenticate token', e);
      return false;
    }
  }
}

// Global singleton instance for easy imports
export const auth = new AuthState();
