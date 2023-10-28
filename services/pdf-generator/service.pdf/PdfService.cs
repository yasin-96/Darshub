using Microsoft.Extensions.Logging;
using QuestPDF.Fluent;
using pdfGen.common.Models.Course;
using pdfGen.service.pdf.Structor;

namespace pdfGen.service.pdf;
public class PdfService
{
    private readonly ILogger<PdfService> _logger;
    private readonly string _filePath;
    //private readonly ParseService _pdParseService;
    
    public PdfService(ILogger<PdfService> logger)
    {
        _filePath = "dhub.pdf";
        _logger = logger;
    }

    public void CreatePdf(CourseWithChapters data)
    {
        try
        {
            var pdfDoc = new PdfStructor(data);
            pdfDoc.GeneratePdf(_filePath);
        }
        catch (Exception e)
        {
            Console.WriteLine(e);
            throw;
        }
    }
}