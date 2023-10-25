using System.Net;

namespace pdfGen.service.parser.Image;
public static class ImageHelper
{
    public static async Task<byte[]> DownloadImageAsync(string imageSrc)
    {
        Console.WriteLine(imageSrc);
        try
        {
            using (var client = new WebClient())
            {
                var resultImage = await client.DownloadDataTaskAsync(imageSrc);
                if (resultImage.Length <= 0)
                {
                    return default;
                }
                
                if (!client.ResponseHeaders["Content-Type"].StartsWith("image/"))
                {
                    return default;
                }

                return resultImage;
            }
        }
        catch (Exception e)
        {
            Console.WriteLine(e);
            return default;
        }
    }
}