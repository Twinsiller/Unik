namespace Pr1
{
    using System.IO;
    internal class Program
    {
        static void readText(string path)
        {
            string line;
            try
            {
                //Pass the file path and file name to the StreamReader constructor
                StreamReader sr = new StreamReader(path);
                //Read the first line of text
                line = sr.ReadLine();
                //Continue to read until you reach end of file
                while (line != null)
                {
                    //write the line to console window
                    Console.WriteLine(line);
                    //Read the next line
                    line = sr.ReadLine();
                }
                //close the file
                sr.Close();
                Console.ReadLine(); //Ожидание
            }
            catch (Exception e)
            {
                Console.WriteLine("Ошибка чтения файла: " + e.Message);
            }
            finally
            {
                Console.WriteLine("Файл прочитан!!!");
            }
        }

        static void findWord(string path)
        {
            string word, line;
            string[] words;
            Console.Write("Введите слово, которое нужно найти в файле: ");
            word = Console.ReadLine();

            try
            {
                //Pass the file path and file name to the StreamReader constructor
                StreamReader sr = new StreamReader(path);
                //Read the first line of text
                line = sr.ReadLine();
                //Continue to read until you reach end of file
                while (line != null)
                {
                    //write the line to console window
                    words = line.Split([' ', '.', ',', ';', '!', '?', ':', '-', '(', ')', '\'', '\"']); 
                    foreach (string i in words)
                    {
                        Console.Write(i + "-");
                        if (i.ToLower() == word.ToLower())
                        {
                            Console.WriteLine("\nСлово найдено");
                            return;
                        }
                        
                    }
                    Console.WriteLine("\n-");
                    //Read the next line
                    line = sr.ReadLine();
                }
                //close the file
                sr.Close();
                Console.ReadLine(); //Ожидание
            }
            catch (Exception e)
            {
                Console.WriteLine("Ошибка чтения файла: " + e.Message);
            }
            Console.WriteLine("Слово НЕ найдено");
        }
        static void Main(string[] args)
        {
            string path;
            Console.Write("Введите путь текстового файла: ");
            path = Console.ReadLine();
            Console.WriteLine(path);

            readText(path); // 1.1
            findWord(path); // 1.2

        }
    }
}
