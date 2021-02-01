export default function documentReady(fn: () => void) {
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', fn);
    return;
  }

  fn();
}
