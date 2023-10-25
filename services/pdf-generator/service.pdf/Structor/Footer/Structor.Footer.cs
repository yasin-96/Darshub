using QuestPDF.Fluent;
using QuestPDF.Helpers;
using QuestPDF.Infrastructure;

namespace pdfGen.service.pdf.Structor;
public partial class PdfStructor
{
    void ComposeFooter(IContainer container)
    {
        container.AlignCenter().Text(x =>
        {
            x.CurrentPageNumber();
            x.Span(" / ");
            x.TotalPages();
        });
    } 
}