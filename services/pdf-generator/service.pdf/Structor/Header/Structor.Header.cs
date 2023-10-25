using QuestPDF.Fluent;
using QuestPDF.Helpers;
using QuestPDF.Infrastructure;

namespace pdfGen.service.pdf.Structor;
public partial class PdfStructor
{
    private void ComposeHeader(IContainer container)
    {
        container.ShowOnce().Row(row =>
        {
            row.RelativeItem().Column(column =>
            {
                column.Item().Text($"#{_pdfData.Id}")
                    .Style(TextStyle.Default.FontSize(16).SemiBold().FontColor(Colors.Blue.Medium));
                column.Item().Text($"#{_pdfData.Name}")
                    .Style(TextStyle.Default.FontSize(14).SemiBold().FontColor(Colors.Blue.Medium));

                column.Item().Text(text =>
                {
                    text.Span("Last update: ").SemiBold();
                    text.Span($"{_pdfData.LastUpdate:d}");
                });
                
                column.Item().Text(text =>
                {
                    text.Span("Release date: ").SemiBold();
                    text.Span($"{_pdfData.Released:d}");
                });
                column.Item().Text(text =>
                {
                    text.Span("Duration Time: ").SemiBold();
                    text.Span($"{_pdfData.Duration}");
                });
                column.Item().Text(text =>
                {
                    text.Span("Level: ").SemiBold();
                    text.Span($"{_pdfData.Level}");
                });
            });
        });
    } 
}