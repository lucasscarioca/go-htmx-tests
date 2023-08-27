console.log('hello')

if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  document.documentElement.classList.add('dark')
  localStorage.theme = 'dark'
  setTimeout(() => {
    const themeIcon = document.getElementById('theme')
    themeIcon?.classList?.replace('ph-moon', 'ph-sun')
  }, 200)
} else {
  document.documentElement.classList.remove('dark')
  localStorage.theme = 'light'
  setTimeout(() => {
    const themeIcon = document.getElementById('theme')
    themeIcon?.classList?.replace('ph-sun', 'ph-moon')
  }, 200)
}

function toggleTheme() {
  const themeIcon = document.getElementById('theme')
  if (localStorage.theme === 'dark') {
    document.documentElement.classList.remove('dark')
    localStorage.theme = 'light'
    themeIcon?.classList?.replace('ph-sun', 'ph-moon')
  } else {
    document.documentElement.classList.add('dark')
    localStorage.theme = 'dark'
    themeIcon?.classList?.replace('ph-moon', 'ph-sun')
  }
}

function menu() {
  const list = document.querySelector('ul')
  const menu = document.getElementById('menu')
  menu.classList.contains('ph-list')
    ? (menu.classList.replace('ph-list', 'ph-x'), list.classList.remove('hidden'), list.classList.add('h-52'), list.classList.add('opacity-100'))
    : (menu.classList.replace('ph-x', 'ph-list'), list.classList.remove('h-52'), list.classList.remove('opacity-100'), list.classList.add('hidden'))
}