using HtmlAgilityPack;

namespace pdfGen.service.parser.HTML;
public static class HtmlParser
{
    public static HtmlDocument? ParseRawHtmlContent(string htmlContent)
    {
        try
        {
            var parsedHtmlStructor =  new HtmlDocument();
            parsedHtmlStructor.LoadHtml(htmlContent);

            if (parsedHtmlStructor.DocumentNode is null)
            {
                
            }

            return parsedHtmlStructor;
        }
        catch (Exception e)
        {
            Console.WriteLine(e);
            return null;
        }
    }
}