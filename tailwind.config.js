/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{astro,html,js}'],
  theme: {
    extend: {
	fontFamily: {
		'serif': ['Playfair Display', 'serif'],
	}
    },
  },
  plugins: [
	require('@tailwindcss/typography'),
  ],
}

