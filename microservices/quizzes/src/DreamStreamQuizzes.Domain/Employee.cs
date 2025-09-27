namespace DreamStreamQuizzes.Domain;

[Table("employees", Schema = "quizzes")]
public class Employee
{
    [Column("id", TypeName = "uuid")]
    [Key]
    public Guid EmployeeId { get; set; }

    [Column("created_at", TypeName = "timestamptz")]
    public DateTimeOffset CreatedAt { get; set; }

    [Column("edited_at", TypeName = "timestamptz")]
    public DateTimeOffset EditedAt { get; set; }

    [Column("first_name", TypeName = "text")]
    public string FirstName { get; set; }

    [Column("second_name", TypeName = "text")]
    public string LastName { get; set; }
}