/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{astro,html,js}'],
  theme: {
    extend: {},
  },
  plugins: [
	require('@tailwindcss/typography'),
  ],
}

