/*
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
    let validString = s.toLowerCase().replace(/ /g, ""); // Удаляем пробелы

    // Перебирает каждый символ в строке
    for (let i = 0; i < validString.length; i++) {
        let currentChar = validString[i];
        // Проверяет, является ли символ буквой или цифрой
        if (!('a' <= currentChar && currentChar <= 'z' || '0' <= currentChar && currentChar <= '9')) {
            // Если символ не буква и не цифра, удаляем его
            validString = validString.substring(0, i) + validString.substring(i + 1);
            i--; // Сдвигаем индекс назад, чтобы проверить следующий символ
        }
    }

    let reverString = "";
    for (let i = validString.length - 1; i >= 0; i--) {
        reverString += validString[i];
    }

    if (validString === reverString) {
        return true;
    }
    return false;
};