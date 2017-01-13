
var fizzbuzzArray = [
    1,
    2,
    "fizz",
    4,
    "buzz",
    "fizz",
    7,
    8,
    "fizz",
    "buzz",
    11,
    "fizz",
    13,
    14,
    "fizzbuzz",
]

for(var cnt = 0; cnt < 100;) {
  for(var num = 0; num < fizzbuzzArray.length && cnt < 100; num++, cnt++) {
    process.stdout.write(fizzbuzzArray[num] + " ")
  }
  for(var num = 0; num < fizzbuzzArray.length; num++) {
    if(isNaN(fizzbuzzArray[num]) === false) {
      fizzbuzzArray[num] += 15
    }
  }
}
