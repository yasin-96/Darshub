using System.ComponentModel.DataAnnotations;

namespace service.mongo.settings;

public class MongoSettings: IMongoSettings
{
    public string URL { get; set; }
    
    public string Database { get; set; }
}

public interface IMongoSettings
{
    [Required]
    public string URL { get; set; }
    
    [Required]
    public string Database { get; set; } 
}