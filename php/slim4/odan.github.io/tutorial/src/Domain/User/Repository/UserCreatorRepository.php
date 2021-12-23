<?php

namespace App\Domain\User\Repository;

use PDO;

/**
 * Repository.
 */
final class UserCreatorRepository
{
    /**
     * @var PDO The database connection
     */
    private $connection;

    /**
     * Constructor.
     *
     * @param PDO $connection The database connection
     */
    public function __construct(PDO $connection)
    {
        $this->connection = $connection;
    }

    /**
     * Insert user row.
     *
     * @param array $user The user
     *
     * @return int The new ID
     */
    public function insertUser(array $user): int
    {
        $row = [
            'username' => $user['username'],
            'first_name' => $user['first_name'],
            'last_name' => $user['last_name'],
            'email' => $user['email'],
        ];

        $sql = "INSERT INTO users (username, first_name, last_name, email) 
                VALUES (:username, :first_name, :last_name, :email);";

        $stmt = $this->connection->prepare($sql);

        $stmt->bindParam('username', $row['username']);
        $stmt->bindParam('first_name', $row['first_name']);
        $stmt->bindParam('last_name', $row['last_name']);
        $stmt->bindParam('email', $row['email']);

        $stmt->execute();

        return (int)$this->connection->lastInsertId();
    }
}
