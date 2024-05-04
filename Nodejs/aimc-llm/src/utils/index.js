export function obj2Url(data) {
  let formData = ''
  for (const it in data) {
    formData += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
  }
  return formData
}
