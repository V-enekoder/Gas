// Este componente de React permite a los usuarios crear un nuevo pedido de bombonas de gas, seleccionando tipos de cilindros y cantidades, para luego proceder al pago.

import React, { useState, useMemo, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { FaTrash, FaPlusSquare, FaInfoCircle } from 'react-icons/fa';
import '../Styles/dashboard.css';
import { orderService, cylinderService } from '../../services/ApiServices';
import { useAuth } from '../../context/AuthContext';


const CreateOrder = () => {
    const navigate = useNavigate();
    const { user } = useAuth();

    const [cylinderTypes, setCylinderTypes] = useState([]);
    const [orderDetails, setOrderDetails] = useState([]);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState(null);

    const [selectedCylinderId, setSelectedCylinderId] = useState('');
    const [quantity, setQuantity] = useState(1);

    useEffect(() => {
        const fetchCylinderTypes = async () => {
            try {
                const response = await cylinderService.getAllTypes();
                const availableCylinders = response.data.filter(c => c.disponible);
                setCylinderTypes(availableCylinders);

                if (availableCylinders.length > 0) {
                    setSelectedCylinderId(availableCylinders[0].id);
                }
            } catch (err) {
                console.error("Error al obtener los tipos de cilindros:", err);
                setError("No se pudieron cargar los productos. Inténtalo de nuevo.");
            }
        };
        fetchCylinderTypes();
    }, []);

    const selectedCylinderObject = useMemo(() => {
        if (!selectedCylinderId) return null;
        return cylinderTypes.find(c => c.id === parseInt(selectedCylinderId, 10));
    }, [selectedCylinderId, cylinderTypes]);


    const handleAddItem = (e) => {
        e.preventDefault();
        const cylinder = selectedCylinderObject;

        if (!cylinder || quantity <= 0) {
            alert("Por favor, selecciona un producto y una cantidad válida.");
            return;
        };

        const newItem = {
            type_cylinder_id: cylinder.id,
            name: cylinder.name,
            price: cylinder.price,
            quantity: quantity,
        };
        
        const existingItemIndex = orderDetails.findIndex(item => item.type_cylinder_id === newItem.type_cylinder_id);
        if (existingItemIndex > -1) {
            const updatedDetails = [...orderDetails];
            updatedDetails[existingItemIndex].quantity += newItem.quantity;
            setOrderDetails(updatedDetails);
        } else {
            setOrderDetails([...orderDetails, newItem]);
        }
    };
    
    const handleRemoveItem = (cylinderId) => {
        setOrderDetails(orderDetails.filter(item => item.type_cylinder_id !== cylinderId));
    };

    const totalPrice = useMemo(() => {
        return orderDetails.reduce((total, item) => total + (item.price * item.quantity), 0);
    }, [orderDetails]);

    const handleCreateOrder = async () => {
        if (orderDetails.length === 0) {
            alert("Debes añadir al menos un producto a tu orden.");
            return;
        }
        setIsLoading(true);
        setError(null);

        const orderPayload = {
            user_id: user.id,
            order_details: orderDetails.map(item => ({
                type_cylinder_id: item.type_cylinder_id,
                quantity: item.quantity,
            })),
        };

        try {
            console.log("Enviando al backend:", orderPayload);
            const response = await orderService.create(orderPayload);
            const createdOrder = response.data;
            
            alert(`¡Pedido creado con éxito! ID: ${createdOrder.id}`);
            navigate(`/orders/${createdOrder.id}/payment`);

        } catch (err) {
            const errorMessage = err.response?.data?.error || "No se pudo crear la orden.";
            setError(errorMessage);
            console.error("Error al crear la orden:", err.response);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="dashboard-content">
            <header className="dashboard-header">
                <h1>Crear Nuevo Pedido</h1>
                <p>Añade los productos y las cantidades que necesitas.</p>
            </header>

            <div className="dashboard-card form-card">
                <h3>Añadir Producto al Pedido</h3>
                <form onSubmit={handleAddItem} className="add-item-form">
                    <div className="form-group">
                        <label>Tipo de Bombona</label>
                        <select
                            value={selectedCylinderId}
                            onChange={(e) => setSelectedCylinderId(e.target.value)}
                            disabled={cylinderTypes.length === 0}
                            className="form-select"
                        >
                            {cylinderTypes.length === 0 ? (
                                <option>Cargando o no hay productos...</option>
                            ) : (
                                cylinderTypes.map((cylinder) => (
                                    <option key={cylinder.id} value={cylinder.id}>
                                        {cylinder.name} - ${cylinder.price.toFixed(2)}
                                    </option>
                                ))
                            )}
                        </select>
                    </div>
                    <div className="form-group" >
                        <label>Cantidad</label>
                        <input
                            type="number"
                            min="1"
                            value={quantity}
                            onChange={(e) => setQuantity(parseInt(e.target.value, 10) || 1)}
                        />
                    </div>
                    <button type="submit" className="btn-add-item" style={{  width: '180px', marginLeft:'35px'}} disabled={cylinderTypes.length === 0}>
                         Añadir
                    </button>
                </form>

                {selectedCylinderObject && (
                    <div className="cylinder-description">
                        <FaInfoCircle />
                        <p>{selectedCylinderObject.description}</p>
                    </div>
                )}
            </div>

            <div className="dashboard-card summary-card order-summary">
                <h3>Resumen del Pedido</h3>
                {orderDetails.length > 0 ? (
                    <>
                        <table className="order-details-table">
                            <thead>
                                <tr>
                                    <th>Producto</th>
                                    <th>Cantidad</th>
                                    <th>Precio Unit.</th>
                                    <th>Subtotal</th>
                                    <th>Acción</th>
                                </tr>
                            </thead>
                            <tbody>
                                {orderDetails.map((item) => (
                                    <tr key={item.type_cylinder_id}>
                                        <td>{item.name}</td>
                                        <td>{item.quantity}</td>
                                        <td>${item.price.toFixed(2)}</td>
                                        <td>${(item.price * item.quantity).toFixed(2)}</td>
                                        <td>
                                            <button
                                                onClick={() => handleRemoveItem(item.type_cylinder_id)}
                                                className="btn-remove-item"
                                                title="Eliminar item"
                                            >
                                                <FaTrash />
                                            </button>
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>

                        <div className="order-total">
                            <h4>Total del Pedido:</h4>
                            <span>${totalPrice.toFixed(2)}</span>
                        </div>
                        
                        {error && <p className="error-message">{error}</p>}
                        
                        <button
                            onClick={handleCreateOrder}
                            className="btn-main-action"
                            disabled={isLoading}
                        >
                            {isLoading ? 'Creando Pedido...' : 'Proceder al Pago'}
                        </button>
                    </>
                ) : (
                    <p>Tu pedido está vacío. Añade productos para continuar.</p>
                )}
            </div>
        </div>
    );
};

export default CreateOrder;