namespace pdfGen.common.Models.Base;
public abstract class BaseContent
{
    [BsonElement("_id")]
    public string Id { get; set; }
}