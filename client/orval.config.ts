import { defineConfig } from 'orval';

export default defineConfig({
  openapi: {
    input: {
      target: '../docs/openapi.yml',
    },
    output: {
      mode: 'split',
      target: 'src/shared/api/generated',
      client: 'react-query',
      override: {
        mutator: {
          path: 'src/shared/api/mutator.ts',
          name: 'mutator',
        },
      },
    },
  },
});