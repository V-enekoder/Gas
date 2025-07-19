// Este componente de React renderiza la página principal del dashboard, mostrando un resumen de la actividad reciente del usuario, incluyendo órdenes, entregas y reportes, y proporcionando enlaces rápidos a las funcionalidades clave.

import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { FaPlus, FaEye } from 'react-icons/fa';
import '../Styles/dashboard.css'
import { useAuth } from '../../context/AuthContext';
import { orderService, deliveryService, reportService } from '../../services/ApiServices';
import CompleteProfileModal from '../components/CompleteProfileModal';


const Dashboard = () => {
  const { user } = useAuth();

  const [recentOrders, setRecentOrders] = useState([]);
  const [recentDeliveries, setRecentDeliveries] = useState([]);
  const [recentReports, setRecentReports] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchDashboardData = async () => {
      if (!user?.id) {
        setIsLoading(false);
        return;
      }
      
      setIsLoading(true);

      try {
        const [ordersRes, deliveriesRes, reportsRes] = await Promise.all([
          orderService.getByUserId(user.id),
          deliveryService.getByUserId(user.id),
          reportService.getByUserId(user.id)
        ]);

        setRecentOrders(Array.isArray(ordersRes.data) ? ordersRes.data.slice(0, 3) : []);
        setRecentDeliveries(Array.isArray(deliveriesRes.data) ? deliveriesRes.data.slice(0, 3) : []);
        setRecentReports(Array.isArray(reportsRes.data) ? reportsRes.data.slice(0, 3) : []);

      } catch (error) {
        console.error("Error al cargar los datos del dashboard:", error);
        setRecentOrders([]);
        setRecentDeliveries([]);
        setRecentReports([]);
      } finally {
        setIsLoading(false); 
      }
    };

    fetchDashboardData();
  }, [user]);

  const CardLoader = () => <p className="card-loader">Cargando...</p>;

 return (
    <div className="dashboard-content">
      <header className="dashboard-header">
        <h1>Buenos días, {user?.name || 'Usuario'}</h1>
        <p>Aquí tienes un resumen de tu actividad.</p>
      </header>
      <div className="dashboard-grid">
        <Link to="/orders/new" className="dashboard-card action-card">
          <div className="action-card-content">
            <FaPlus className="action-icon" />
            <h2>Crear un Pedido</h2>
            <p>Pide tus bombonas de gas aquí.</p>
          </div>
        </Link>

        <div className="dashboard-card summary-card">
          <h3>Órdenes Recientes</h3>
          {isLoading ? <CardLoader /> : (
            recentOrders.length > 0 ? (
              <ul>
                {recentOrders.map(order => (
                  <li key={order.id}>
                    Pedido ORD-{order.id}
                    <span className={`status-badge-sm status-${order.order_state.name.toLowerCase()}`}>
                      {order.order_state.name}
                    </span>
                  </li>
                ))}
              </ul>
            ) : <p className="card-empty-text">No tienes órdenes recientes.</p>
          )}
          <Link to="/orders" className="view-all-link">Ver todas <FaEye /></Link>
        </div>

        <div className="dashboard-card summary-card">
          <h3>Últimas Entregas</h3>
          {isLoading ? <CardLoader /> : (
            recentDeliveries.length > 0 ? (
              <ul>
                {recentDeliveries.map(delivery => (
                  <li key={delivery.id}>
                    Entrega DEL-{delivery.id}
                  </li>
                ))}
              </ul>
            ) : <p className="card-empty-text">No tienes entregas programadas.</p>
          )}
          <Link to="/deliveries" className="view-all-link">Ver historial <FaEye /></Link>
        </div>

        <div className="dashboard-card summary-card">
          <h3>Tus Reportes</h3>
          {isLoading ? <CardLoader /> : (
            recentReports.length > 0 ? (
              <ul>
                {recentReports.map(report => (
                  <li key={report.id} >
                    Reporte REP-{report.id}
                    <span className={`status-badge-sm status-${report.report_state.name.toLowerCase().replace(' ', '-')}`}>
                      {report.report_state.name}
                    </span>
                  </li>
                ))}
              </ul>
            ) : <p className="card-empty-text">No tienes reportes activos.</p>
          )}
          <Link to="/reports" className="view-all-link">Ver todos <FaEye /></Link>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;