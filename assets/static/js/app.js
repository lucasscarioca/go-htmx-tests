console.log('hello')

if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  document.documentElement.classList.add('dark')
  localStorage.theme = 'dark'
} else {
  document.documentElement.classList.remove('dark')
  localStorage.theme = 'light'
}

function toggleTheme() {
  if (localStorage.theme === 'dark') {
    document.documentElement.classList.remove('dark')
    localStorage.theme = 'light'
  } else {
    document.documentElement.classList.add('dark')
    localStorage.theme = 'dark'
  }
}

function menu(e) {
  let list = document.querySelector('ul')
  e.name === 'menu'
    ? (e.name = 'close', list.classList.add('top-[50px]'), list.classList.add('opacity-100'), list.classList.remove('z-[-1]'))
    : (e.name = 'menu', list.classList.remove('top-[50px]'), list.classList.remove('opacity-100'), list.classList.add('z-[-1]'))
}