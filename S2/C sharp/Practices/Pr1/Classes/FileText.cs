using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Xml.Linq;

namespace Pr1.Classes
{
    internal class FileText
    {
        private string path; // Путь к файлу
        public string Path // Геттер и Сеттер
        {
            get
            {
                return path;    // возвращаем значение свойства
            }
            set
            {
                path = value;   // устанавливаем новое значение свойства
            }
        }
        public FileText(string _path)
        {
            FileInfo fi = new FileInfo(_path);
            if (fi.Exists) path = _path;
            else
            {
                path = "C:\\Users\\boldi\\Desktop\\Something\\Unik\\S2\\C sharp\\Practices\\Pr1\\Check.txt";
                Console.WriteLine("Такого файла не существует. Путь заменён на дефолтный: \n" + path);
            }
        }

        /// <summary>
        /// Чтение текста из файла
        /// </summary>
        /// <param name="path">Путь к файлу</param>
        public void readText()
        {
            try
            {
                StreamReader sr = new StreamReader(path); // Объект для чтения файла
                string line = sr.ReadLine(); // Чтение первой строки файла

                while (line != null) // Пока не закончатся строки
                {
                    Console.WriteLine(line); // Вывод строки
                    line = sr.ReadLine(); // Чтение следующей строки
                }

                sr.Close(); // Закрытие файла
                Console.ReadKey(); // Ожидание
            }
            catch (Exception e) // Вывод ошибки, если есть
            {
                Console.WriteLine("Ошибка чтения файла: " + e.Message);
                return;
            }
            Console.WriteLine("Файл прочитан!!!");
        }


        /// <summary>
        /// Поиск слова
        /// </summary>
        /// <param name="path">Путь к файлу</param>
        public void findWord()
        {
            try
            {
                StreamReader sr = new StreamReader(path); // Объект для чтения файла
                string line = sr.ReadLine(); // Чтение первой строки файла

                string word; // Переменная для слова, которое ищем
                string[] words; // Переменная для слов из строки
                Console.Write("Введите слово, которое нужно найти в файле: ");
                word = Console.ReadLine();

                while (line != null) // Пока не закончатся строки
                {
                    words = line.Split([' ', '.', ',', ';', '!', '?', ':', '-', '(', ')', '\'', '\"']); // Разделяем строку на слова
                    foreach (string i in words)
                    {
                        if (i.ToLower() == word.ToLower()) // Сравнение слов
                        {
                            Console.WriteLine("\nСлово найдено");
                            return;
                        }
                    }
                    line = sr.ReadLine();
                }
                sr.Close(); // Закрытие файла
                Console.ReadLine(); //Ожидание
            }
            catch (Exception e) // Вывод ошибки, если есть
            {
                Console.WriteLine("Ошибка чтения файла: " + e.Message);
                return;
            }
            Console.WriteLine("Слово НЕ найдено");
        }

        /// <summary>
        /// Задание 4.1 - Подсчёт количества определённого слова в файле
        /// </summary>
        /// <param name="path">Путь к файлу</param>
        /// <returns>Количество определённого слова</returns>
        public int countWord()
        {
            int count = 0;
            try
            {
                StreamReader sr = new StreamReader(path); // Объект для чтения файла
                string? line = sr.ReadLine(); // Чтение первой строки файла

                string word; // Переменная для слова, которое ищем
                string[] words; // Переменная для слов из строки
                Console.Write("Введите слово, которое нужно найти в файле: ");
                word = Console.ReadLine();

                while (line != null) // Пока не закончатся строки 
                {
                    words = line.Split([' ', '.', ',', ';', '!', '?', ':', '-', '(', ')', '\'', '\"']); // Разделяем строку на слова
                    foreach (string i in words) if (i.ToLower() == word.ToLower()) count += 1; // Поиск слова в строке
                    line = sr.ReadLine(); ; // Чтение следующей строки
                }
                sr.Close(); // Закрытие файла
                Console.ReadLine(); // Ожидание
            }
            catch (Exception e) // Вывод ошибки, если есть
            {
                Console.WriteLine("Ошибка чтения файла: " + e.Message);
            }
            return count;
        }

        /// <summary>
        /// Задание 4.1 - Подсчёт общего количества слов в файле
        /// </summary>
        /// <param name="path">Путь к файлу</param>
        /// <returns>Общее количество слов в файле</returns>
        public int countAllWords()
        {
            int count = 0; // Переменная "Количество слов"
            try
            {
                StreamReader sr = new StreamReader(path); // Объект для чтения файла
                string line = sr.ReadLine(); // Чтение первой строки файла
                string[] words; // Переменная для слов из строки

                while (line != null) // Пока не закончатся строки 
                {
                    words = line.Split([' ', '.', ',', ';', '!', '?', ':', '-', '(', ')', '\'', '\"']); // Разделяем строку на слова
                    foreach (string i in words) if (i != "") count += 1; // Поиск слов
                    line = sr.ReadLine(); // Чтение следующей строки
                }
                sr.Close(); // Закрытие файла
                Console.ReadLine(); // Ожидание
            }
            catch (Exception e) // Вывод в случае ошибки
            {
                Console.WriteLine("Ошибка чтения файла: " + e.Message);
            }
            return count;
        }
    }
}
