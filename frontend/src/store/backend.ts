import { createSignal, onMount, onCleanup } from 'solid-js';
import { api } from '../api/client';

const [_backendAvailable, setBackendAvailable] = createSignal<boolean>(false);
const [_gitSha, setGitSha] = createSignal<string>('');
const [_buildTime, setBuildTime] = createSignal<string>('');

export function backendAvailable(): boolean {
  return _backendAvailable();
}

export function gitSha(): string {
  return _gitSha();
}

export function buildTime(): string {
  return _buildTime();
}

/**
 * Initializes polling for backend status and metadata.
 * Call this inside a root component to share status across the app.
 */
export function initBackendStatusPolling(intervalMs: number = 30000): void {
  onMount(() => {
    async function checkBackend(): Promise<void> {
      try {
        const statusResponse = await api.GET('/v1/status');
        const statusData = statusResponse.data;
        if (statusData?.status === 'ok') {
          setBackendAvailable(true);
        } else {
          setBackendAvailable(false);
        }
      } catch (err) {
        setBackendAvailable(false);
        if (import.meta.env.DEV) {
          console.error('[backend] status check failed:', err);
        }
      }
    }

    async function fetchMetadata(): Promise<void> {
      try {
        const metadataResponse = await api.GET('/v1/metadata');
        const metadata = metadataResponse.data;
        if (metadata) {
          setGitSha(metadata.git_sha);
          setBuildTime(metadata.build_time);
        }
      } catch (err) {
        if (import.meta.env.DEV) {
          console.error('[backend] metadata fetch failed:', err);
        }
      }
    }

    void checkBackend();
    void fetchMetadata();
    const timer = setInterval(() => void checkBackend(), intervalMs);
    onCleanup(() => clearInterval(timer));
  });
}
