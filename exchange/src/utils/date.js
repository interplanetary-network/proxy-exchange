export function toLocal(seconds) {
  const date = new Date(seconds * 1000);
  return `${date.getFullYear()}-${padding(date.getMonth() + 1)}-${padding(date.getDate())}T${padding(date.getHours())}:${padding(date.getMinutes())}`;
}

export function toSeconds(string) {
  const date = new Date(string);
  return date.getTime() / 1000;
}

export function isPast(seconds) {
  return new Date(seconds * 1000) < new Date()
}

function padding(string) {
  return ('0' + string).slice(-2);
}
