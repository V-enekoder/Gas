// Este componente de React muestra una lista de reportes generados por el usuario, permitiendo ver su historial y crear nuevos reportes.

import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { FaEye, FaPlus, FaFileMedicalAlt } from 'react-icons/fa';
import '../Styles/dashboard.css';
import { useAuth } from '../../context/AuthContext'
import { reportService } from '../../services/ApiServices';


const Reports = () => {
  const [reports, setReports] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const { user } = useAuth();

  useEffect(() => {
    const fetchUserReports = async () => {
      if (!user?.id) {
        setIsLoading(false);
        return;
      }
      setIsLoading(true);
      try {
        const response = await reportService.getByUserId(user.id);
        if (Array.isArray(response.data)) {
          setReports(response.data);
        } else {
          setReports([]);
        }
      } catch (error) {
        console.error("Error al obtener los reportes:", error);
        setReports([]);
      } finally {
        setIsLoading(false);
      }
    };
    fetchUserReports();
  }, [user]);

  const formatDate = (dateString) => new Date(dateString).toLocaleDateString('es-ES', { day: '2-digit', month: 'long', year: 'numeric' });

  if (isLoading) return <div className="loading-spinner">Cargando tus reportes...</div>;

  return (
    <div className="dashboard-content">
      <header className="dashboard-header-with-action">
        <div>
          <h1>Mis Reportes</h1>
          <p>Historial de todos los reportes que has creado.</p>
        </div>
        <Link to="/reports/new" className="btn-main-action header-btn">
          <FaPlus /> Crear Nuevo Reporte
        </Link>
      </header>
      
      {reports.length === 0 ? (
        <div className="empty-state">
          <FaFileMedicalAlt size={50} style={{ marginBottom: '20px', color: 'var(--fruit-salad-400)' }} />
          <h2>No tienes reportes creados</h2>
          <p>Si tienes algún problema con una entrega, puedes crear un reporte aquí.</p>
        </div>
      ) : (
        <div className="dashboard-card orders-table-card">
          <table className="orders-table">
            <thead>
              <tr>
                <th>ID del Reporte</th>
                <th>Fecha</th>
                <th>Tipo</th>
                <th>Estado</th>
                <th>Acciones</th>
              </tr>
            </thead>
            <tbody>
              {reports.map((report) => (
                <tr key={report.id}>
                  <td><span className="order-id-highlight">REP-{report.id}</span></td>
                  <td>{formatDate(report.date)}</td>
                  <td>{report.report_type.name}</td>
                  <td>
                    <span className={`status-badge status-${report.report_state.name.toLowerCase().replace(' ', '-')}`}>
                      {report.report_state.name}
                    </span>
                  </td>
                  <td>
                    <Link to={`/reports/${report.id}`} className="action-btn">
                      <FaEye /> Ver Detalles
                    </Link>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
};

export default Reports;