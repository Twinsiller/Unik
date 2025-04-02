using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using System.Windows;

namespace Pr2.Services
{
    class FilesAnalisator
    {
        private readonly object _lock = new object();
        private readonly List<FileAnalysis> _results = new List<FileAnalysis>();

        public FilesAnalisator() { }

        public async Task<List<FileAnalysis>> AnalyzeFilesAsync(string[] filePaths)
        {
            _results.Clear();
            var tasks = filePaths.Select(ProcessFileAsync).ToArray();
            await Task.WhenAll(tasks);
            return _results;
        }

        private async Task ProcessFileAsync(string filePath)
        {
            try
            {
                string content = await File.ReadAllTextAsync(filePath);
                int wordCount = content.Split(new[] { ' ', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries).Length;
                int charCount = content.Length;

                var analysis = new FileAnalysis(filePath, wordCount, charCount);

                lock (_lock)
                {
                    _results.Add(analysis);
                }
            }
            catch (Exception ex)
            {
                throw new IOException($"Ошибка обработки файла {filePath}: {ex.Message}", ex);
            }
        }
    }
}
