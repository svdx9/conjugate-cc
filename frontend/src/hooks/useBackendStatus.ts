import { createSignal, onMount, onCleanup } from 'solid-js';
import { api } from '../api/client';

export interface BackendStatus {
  available: () => boolean;
  gitSha: () => string;
  buildTime: () => string;
}

export function useBackendStatus(intervalMs: number = 30000): BackendStatus {
  const [available, setAvailable] = createSignal(false);
  const [gitSha, setGitSha] = createSignal('');
  const [buildTime, setBuildTime] = createSignal('');

  onMount(() => {
    async function checkBackend(): Promise<void> {
      try {
        const statusResponse = await api.GET('/v1/status');
        const statusData = statusResponse.data;
        if (statusData?.status === 'ok') {
          setAvailable(true);
        } else {
          setAvailable(false);
        }
      } catch {
        setAvailable(false);
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
      } catch {
        // Silently fail metadata
      }
    }

    void checkBackend();
    void fetchMetadata();
    const timer = setInterval(() => void checkBackend(), intervalMs);
    onCleanup(() => clearInterval(timer));
  });

  return { available, gitSha, buildTime };
}
