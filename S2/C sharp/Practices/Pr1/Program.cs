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
            path = "C:\\Users\\boldi\\OneDrive\\Desktop\\Something\\Unik\\S2\\C sharp\\Practices\\Pr1\\Check.txt";//Console.ReadLine();
            //path = Console.ReadLine();
            Console.WriteLine(path);

            

            int number;
            do
            {
                Console.Write("\n\nВыберите команду для действия:" +
                "\n1 - (1.1)Вывод текста;" +
                "\n2 - (1.2)Поиск слова в тексте;" +
                "\n3 - (2.1);" +
                "\n4 - (2.2);" +
                "\n5 - (3.1);" +
                "\n6 - (3.2);" +
                "\n\nВведите число: ");
                number = Console.Read();
                switch (number)
                {
                    case 1:
                        readText(path); // 1.1
                        break;
                    case 2:
                        findWord(path); // 1.2
                        break;
                    default:
                        Console.WriteLine("Exiting program");
                        break;
                }
                Console.WriteLine(number);
                Console.ReadLine();
                //Console.Clear();
            } while (0 < number && number < 3);



        }
    }
}
