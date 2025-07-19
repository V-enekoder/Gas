// Este archivo define el componente de React para la ventana modal que muestra los detalles de un reporte.
import React from 'react';
import { FaTimes, FaClipboardList, FaCalendarAlt, FaFlag, FaFileAlt, FaTruck, FaFileInvoice } from 'react-icons/fa';
import '../Styles/components.css';
const ReportDetailsModal = ({ report, onClose }) => {
  if (!report) return null;

  const formatDate = (dateString) =>
    new Date(dateString).toLocaleDateString('es-ES', {
      day: '2-digit',
      month: 'long',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });

  return (
    <div className="modal-overlay">
      <div className="modal-container order-detail-card">
        <button className="modal-close-btn" onClick={onClose}>
          <FaTimes />
        </button>

        <h2 className="modal-title">Detalles del Reporte <span className="order-id-highlight">REP-{report.id}</span></h2>
        <p className="modal-subtitle">Toda la información sobre el reporte generado.</p>

        {/* Sección Información General */}
        <div className="detail-section">
          <h2><FaClipboardList /> Información General</h2>
          <div className="detail-grid">
            <div className="detail-item">
              <FaCalendarAlt />
              <span>Fecha del Reporte</span>
              <strong>{formatDate(report.date)}</strong>
            </div>
            <div className="detail-item">
              <FaFileAlt />
              <span>Tipo de Reporte</span>
              <strong>{report.report_type.name}</strong>
            </div>
            <div className="detail-item">
              <FaFlag />
              <span>Estado</span>
              <strong>{report.report_state.name}</strong>
            </div>
          </div>
        </div>

        {/* Sección Descripción */}
        <div className="detail-section">
          <h2><FaFileInvoice /> Descripción del Reporte</h2>
          <p className="description-box">{report.description}</p>
        </div>

        {/* Sección de Entrega Asociada */}
        <div className="detail-section">
          <h2><FaTruck /> Entrega Asociada</h2>
          <div className="detail-grid">
            <div className="detail-item">
              <span>ID de Entrega</span>
              <strong>DEL-{report.delivery.id}</strong>
            </div>
            <div className="detail-item">
              <span>ID de Orden</span>
              <strong>ORD-{report.delivery.order_id}</strong>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ReportDetailsModal;