import type { Config } from 'tailwindcss'

export default {
  content: [
    './src/**/*.{js,jsx,ts,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#007bff',
          foreground: '#ffffff',
        },
        secondary: {
          DEFAULT: '#6c757d',
          foreground: '#ffffff',
        },
        destructive: {
          DEFAULT: '#dc3545',
        },
        background: '#ffffff',
        accent: {
          DEFAULT: '#f8f9fa',
          foreground: '#212529',
        },
      },
    },
  },
  plugins: [],
} satisfies Config