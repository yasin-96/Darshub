using pdfGen.common.Models.Course;
using QuestPDF.Fluent;
using QuestPDF.Infrastructure;

namespace pdfGen.service.pdf.Structor;
public partial class PdfStructor: IDocument
{
    private readonly CourseWithChapters _pdfData;

    public PdfStructor(CourseWithChapters pdfData)
    {
        _pdfData = pdfData;
        QuestPDF.Settings.License = LicenseType.Community;
    }
    
    public void Compose(IDocumentContainer container)
    {
        container
            .Page(page =>
            {
                page.Margin(50);
            
                page.Header().Element(ComposeHeader);
                page.Content().Element(ComposeContent);
                page.Footer().Element(ComposeFooter);
            });
    }
}