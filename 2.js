// const addTwoNumbers = (l1, l2) => {
    
//     let newList = new NodeList()
//     let sum
//     let reversValue1 = ""
//     let reverseValue2 = ""

//     l1.forEach(element => {
//         reversValue1 = element.value + reversValue1
//     })

//     l2.forEach(element => {
//         reversValue2 = element.value + reversValue2
//     })

//     sum = parseInt(reversValue1) + parseInt(reverseValue2)


// }





/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    let dummyHead = new ListNode(0); // Создаем фиктивный узел для удобства
    let current = dummyHead; // Текущий узел в результирующем списке
    let carry = 0; // Перенос единицы

    while (l1 || l2 || carry) { // Проходим по обоим спискам, пока в них есть узлы или перенос
        let x = l1 ? l1.val : 0; // Получаем значение из l1, если оно есть, иначе 0
        let y = l2 ? l2.val : 0; // Получаем значение из l2, если оно есть, иначе 0

        let sum = x + y + carry; // Сумма значений узлов и переноса
        carry = Math.floor(sum / 10); // Новый перенос
        current.next = new ListNode(sum % 10); // Создаем новый узел с остатком от деления суммы на 10

        current = current.next; // Переходим к следующему узлу в результирующем списке
        l1 = l1 ? l1.next : null; // Переходим к следующему узлу в l1, если он есть
        l2 = l2 ? l2.next : null; // Переходим к следующему узлу в l2, если он есть
    }

    return dummyHead.next; // Возвращаем следующий узел после фиктивного узла, то есть начало результирующего списка
};
