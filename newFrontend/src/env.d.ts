/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_PROTO: string;
  readonly VITE_WORKER_PORT: string;
  readonly VITE_WORKER_HOST: string;
  readonly VITE_WORKER_URI: string;
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
