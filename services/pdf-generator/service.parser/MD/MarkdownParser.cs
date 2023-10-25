using Markdig;

namespace pdfGen.service.parser.MD;
public static class MarkdownParser
{
    public static string? ParseMarkdownToHtmlContent(string mdContent)
    {
        try
        {
            var pipeline = new MarkdownPipelineBuilder().UseAdvancedExtensions().Build();
            var parsedHtmlContent = Markdown.ToHtml(mdContent, pipeline);

            if (string.IsNullOrEmpty(parsedHtmlContent))
            {
                return string.Empty;
            }

            if (parsedHtmlContent.Length < 0)
            {
                return string.Empty;
            }

            return parsedHtmlContent;
        }
        catch (Exception e)
        {
            Console.WriteLine(e);
            return string.Empty;
        }
    }
}