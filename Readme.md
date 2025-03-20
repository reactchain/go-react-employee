# 🏢 Employee Hierarchy System

## 🚀 Overview
This project provides a **Go backend** that serves employee hierarchy data and a **React frontend** that displays it in a **nested structure**. The system ensures that employees are grouped under their respective managers, sorted alphabetically by last name.

---

## 📂 Project Structure
```
employee-hierarchy/
│── backend/              # Go Backend (Gin Framework)
│   ├── main.go           # Main Go application
│── frontend/             # React Frontend
│   ├── src/
│   │   ├── EmployeeHierarchy.js  # React Component
│   │   ├── App.js        # Main App Entry
│   ├── package.json      # React Dependencies
│── README.md             # Project Documentation
```

---

## 🎯 Features
✅ **Go Backend** - Serves employee data via REST API  
✅ **Predefined Employee List** - Hardcoded employee data in `main.go`  
✅ **Recursive Hierarchy Rendering** - Displays employees nested under managers  
✅ **Sorted Employees** - Reports sorted alphabetically by last name  
✅ **CORS Enabled** - Allows frontend to fetch from backend  
✅ **React Frontend** - Fetches and renders employee hierarchy  

---

## 🔧 Installation & Setup

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/your-repo/employee-hierarchy.git
cd employee-hierarchy
```

---

### 2️⃣ Run the Go Backend
```sh
cd backend
go mod init employee-hierarchy
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go run main.go
```
The server will be running at: **`http://localhost:8080`**

---

### 3️⃣ Run the React Frontend
```sh
cd frontend
npm install
npm start
```
The React app will be available at: **`http://localhost:3000`**

---

## 🔗 API Endpoints
| Method | Endpoint          | Description                 |
|--------|------------------|-----------------------------|
| GET    | `/employees`      | Fetch employee hierarchy   |

### 📌 Example Response:
```json
[
  {
    "name": "Michael Chen",
    "title": "CEO",
    "manager_id": null,
    "reports": [
      {
        "name": "Barrett Glasauer",
        "title": "CTO",
        "manager_id": 2,
        "reports": [
          {
            "name": "Chris Hancock",
            "title": "Engineering Manager",
            "manager_id": 1,
            "reports": [
              {
                "name": "Shrutika Dasgupta",
                "title": "Engineer",
                "manager_id": 8,
                "reports": []
              }
            ]
          }
        ]
      }
    ]
  }
]
```

---

## 🛠 Troubleshooting
### 🚨 CORS Issue?
Ensure the Go backend has **CORS enabled** in `main.go`:
```go
r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
}))
```

---

## 🤝 Contributing
🔹 **Fork** the repository  
🔹 **Create** a feature branch  
🔹 **Commit** changes  
🔹 **Submit** a **pull request** 🎉  

---

## 📜 License
MIT License © 2025 Your Name  
