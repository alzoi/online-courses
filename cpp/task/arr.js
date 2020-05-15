// На входе все числа, кроме одного, имеют пару. Найти число без пары.

// Функция для работы со стандартным вводом.
// https://nodejs.org/api/readline.html
let readline = require('readline');
let rl = readline.createInterface({
  input:    process.stdin,
  output:   process.stdout,
  terminal: false
});

let arr = [];

// Событие - передана строка.
rl.on('line', function(line){    
  // Добавляем значение в массив.
  arr.push(line);
})

// Событие завершение чтения файла стандартного ввода.
rl.on('close', (input) => {
  // Получаем объект со свойствами "число" : "количество повторов".
  var result = arr.reduce(function(acc, el) {
    acc[el] = (acc[el] || 0) + 1;
    return acc;
  }, {});
  
  // Находим свойство объекта со значением 1.
  console.log(Object.keys(result).find(key => result[key] === 1));
});
