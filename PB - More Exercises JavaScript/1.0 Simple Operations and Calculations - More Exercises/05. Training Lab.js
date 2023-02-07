([w, h]) => (w * 100 / 120 | 0) * (((h * 100) - 100) / 70 | 0) - 3

//----------------------------------------------------------------------------or (2)
function trainingLab([width, len]) {
  let rows = Math.trunc(width * 100 / 120)
  let desk = Math.trunc(((len * 100) - 100) / 70)
  let allPlace = rows * desk - 3
  return allPlace
}
//console.log(trainingLab([15, 8.9]))
