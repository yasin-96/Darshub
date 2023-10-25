using HTMLQuestPDF.Extensions;
using pdfGen.service.parser.Image;
using pdfGen.service.parser.MD;
using QuestPDF.Fluent;
using QuestPDF.Helpers;
using QuestPDF.Infrastructure;

namespace pdfGen.service.pdf.Structor;
public partial class PdfStructor
{
    private void ComposeContent(IContainer container)
    {
        container.Column(colChapter =>
        {
            colChapter.Item().AlignCenter().Text(_pdfData.Name).Style(new TextStyle().Bold().Light().FontSize(14));
            colChapter.Item().Text("Description:").Style(new TextStyle().Bold().Light().FontSize(10));
            colChapter.Item().Text(_pdfData.Description).Style(new TextStyle().Bold().Light().FontSize(10));
            colChapter.Item().PaddingVertical(5).LineHorizontal(1).LineColor(Colors.Grey.Medium);
            for (var chapPos = 0; chapPos < _pdfData.Content.Count; chapPos++)
            {
                var chapter = _pdfData.Content[chapPos];

                colChapter.Item().PaddingTop(10).Text(chapter.Title).Style(new TextStyle()
                    .Underline()
                    .FontSize(16)
                    .Italic()
                );  
                colChapter.Item().PaddingTop(8).HTML(h => h.SetHtml(MarkdownParser.ParseMarkdownToHtmlContent(chapter.Description)));
                    
                colChapter.Item().PaddingTop(5).Column(colSubChapter =>
                {
                    for (var subPos = 0; subPos < chapter.SubChapterReferences.Count; subPos++)
                    {
                        var sub = chapter.SubChapterReferences[subPos];
                        var parsedSubChapter = MarkdownParser.ParseMarkdownToHtmlContent(sub.Content);
                            
                        colSubChapter.Item().Text($"{subPos +1}.) {sub.Title}").Style(new TextStyle()
                                .FontSize(14)
                                .Italic()
                            ); 
                        colSubChapter.Item().HTML(handler =>
                        {
                            handler.OverloadImgReceivingFunc( (x)=> ImageHelper.DownloadImageAsync(x).Result);
                            //TODO: Change html style after Test
                            handler.SetHtml(parsedSubChapter);
                        });
                        
                        colSubChapter.Item().PaddingTop(5);
                    }
                });
            }
        });
    }
}