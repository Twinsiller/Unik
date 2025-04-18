using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Pr2.Services
{
    public class FileAnalysis
    {
        public string _filePath { get;}
        public string _fileName { get; }
        public int _countWords { get; }
        public int _countChars { get; }

        public FileAnalysis(string filePath, int countWords, int countChars)
        {
            _filePath = filePath;
            _fileName = GetFileName(filePath);
            _countWords = countWords;
            _countChars = countChars;
        }

        public string GetFileName(string filePath)
        {
            return Path.GetFileName(filePath);
        }
    }
}
