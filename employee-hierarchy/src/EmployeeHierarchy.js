import React, { useEffect, useState } from "react";

const EmployeeHierarchy = () => {
  const [hierarchy, setHierarchy] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/employees")
      .then((res) => res.json())
      .then((data) => setHierarchy(data))
      .catch((err) => console.error("Error fetching data:", err));
  }, []);

  // Recursive function to render nested employees
  const renderEmployees = (employees, indent = 0) => {
    return (
      <ul>
        {employees.map((emp,i) => (
          <li key={i} style={{ marginLeft: indent * 20 }}>
            <strong>{emp.title}:</strong> {emp.name}
            {emp.reports.length > 0 && renderEmployees(emp.reports, indent + 1)}
          </li>
        ))}
      </ul>
    );
  };

  return (
    <div>
      <h2>Employee Hierarchy</h2>
      {renderEmployees(hierarchy)}
    </div>
  );
};

export default EmployeeHierarchy;
