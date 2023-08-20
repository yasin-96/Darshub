using pdfGen.common.Models.Base;

namespace pdfGen.common.Models.Course;
public class Chapter : BaseContent
{
    [BsonElement("name")]
    public string Name { get; set; }

    [BsonElement("description")]
    public string Description { get; set; }

    [BsonElement("skills")]
    public string Skills { get; set; }

    [BsonElement("subchapters")]
    public List<string> SubChapterReferences { get; set; }
}

public class SubChapter : BaseContent 
{
    [BsonElement("content")]
    public byte[] Content { get; set; }

    [BsonElement("listing")]
    public void Listing { get; set; }
}