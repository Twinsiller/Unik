
namespace Pr1
{
    #include <clocale>
    using Pr1.Classes;
    using System.IO;
    using static System.Runtime.InteropServices.JavaScript.JSType;

    internal class Program
    {
        static void Main(string[] args)
        {
            string path; // Переменная "Путь к файлу"
            //Console.Write("Введите путь текстового файла: ");
            path = @"C:\Users\boldi\Desktop\Something\Unik\S2\C sharp\Practices\Pr1\Check.txt";
            //path = "Check.txt";
            //path = Console.ReadLine()

            FileText ft = new FileText(path);

            string number;
            do
            {
                Console.WriteLine("Ваш путь: " + ft.Path + "\n");
                Console.Write("Выберите команду для действия:" +
                "\n1 - (1.1) Вывод текста;" +
                "\n2 - (1.2) Поиск слова в тексте;" +
                "\n3 - (2.1) ;" +
                "\n4 - (2.2) ;" +
                "\n5 - (3.1) Количество определённого слова в файле;" +
                "\n6 - (3.2) ;" +
                "\n7 - (4.1) Общее количество слов в файле;" +
                "\n8 - (4.2) ;" +
                "\nЛюбое другое число или символы это выход из программы" +
                "\n\nВведите число: ");
                number = Console.ReadLine();
                if (int.TryParse(number, out var x)) // Проверка вводимого значения
                    switch (x)
                    {
                        case 1:
                            ft.readText(); // 1.1
                            break;
                        case 2:
                            ft.findWord(); // 1.2
                            break;
                        case 3:
                            
                            break;
                        case 4:

                            break;
                        case 5:
                            Console.WriteLine("Количество таких слов: " + ft.countWord());
                            break;
                        case 6:

                            break;
                        case 7:
                            Console.WriteLine("Количество слов в файле: " + ft.countAllWords());
                            break;
                        case 8:

                            break;
                        default:
                            Console.WriteLine("Exiting program"); // Завершение программы
                            break;
                    }
                else
                {
                    Console.WriteLine("Exiting program"); // Завершение программы
                    break;
                }
                Console.WriteLine("Выполнена команда под номером " + number); // Что выполнили
                Console.ReadKey(); // Ожидание
                Console.Clear(); // Очистка консоли
            } while (true);
        }
    }
}
