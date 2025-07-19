import { Routes, Route } from 'react-router-dom';

import Login from '../page/form/Login';
import Register from '../page/form/Register';
import Dashboard from '../page/home/Dashboard';

import DashboardLayout from '../page/components/DashboardLayout'; 
import CreateOrder from '../page/order/CreateOrder';
import Payment from '../page/order/Payment';
import OrdersPlaced from '../page/order/OrdersPlaced';
import Deliveries from '../page/deliveries/Deliveries';
import CreateReport from '../page/components/CreateReport';
import Reports  from '../page/reports/Reports';
import Home from '../page/Home/Home';
import OrderDetail from '../page/components/OrderDetail';
import DeliveryReceipt from '../page/components/DeliveryReceipt';

export const AppRouter = () => {
    return (
        <Routes>
            <Route path='/' element={<Home/>}/>
            <Route path='/login' element={<Login/>}/>
            <Route path='/register' element={<Register/>}/>

            <Route element={<DashboardLayout />}>
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/orders/new" element={<CreateOrder />} />
                <Route path="/orders/:orderId/payment" element={<Payment />} />
                <Route path="/orders" element={<OrdersPlaced />} />
                <Route path="/deliveries" element={<Deliveries />} />
                <Route path="/reports" element={<Reports />} />
                <Route path="/reports/new" element={<CreateReport />} />
                 <Route path="/orders/:orderId" element={<OrderDetail />} />
                 <Route path="/deliveries/:id" element={<DeliveryReceipt />} />
            </Route>
        </Routes>
    );
};


export default AppRouter;