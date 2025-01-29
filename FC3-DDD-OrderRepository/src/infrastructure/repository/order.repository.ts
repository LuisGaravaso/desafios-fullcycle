

//Implements
// export default interface RepositoryInterface<T> {
//     create(entity: T): Promise<void>;
//     update(entity: T): Promise<void>;
//     find(id: string): Promise<T>;
//     findAll(): Promise<T[]>;
// }

import Order from "../../domain/entity/order";
import OrderModel from "../db/sequelize/model/order.model";
import OrderRepositoryInterface from "../../domain/repository/order-repository-interface";
import OrderItemModel from "../db/sequelize/model/order_item.model";
import OrderItem from "../../domain/entity/order_item";

export default class OrderRepository implements OrderRepositoryInterface {
    async create(entity: Order): Promise<void> {
        //Created OrderModel
        await OrderModel.create({
            id: entity.id,
            customer_id: entity.customerId,
            total: entity.total(),
            items: entity.items.map((item) => ({
                id: item.id,
                name: item.name,
                price: item.price,
                product_id: item.productId,
                quantity: item.quantity,
            })),
        },
        {
            include: [{model: OrderItemModel, as: 'items'}],
        });
    }

    async update(entity: Order): Promise<void> {
        const transaction = await OrderModel.sequelize.transaction();
    
        try {
            // Update the OrderModel
            await OrderModel.update(
                {
                    customer_id: entity.customerId,
                    total: entity.total(),
                },
                {
                    where: { id: entity.id },
                    transaction,
                }
            );
    
            // Fetch existing items in the database for this order
            const existingItems = await OrderItemModel.findAll({
                where: { order_id: entity.id },
                transaction,
            });
    
            // Find items to delete (exist in DB but not in entity.items)
            const itemsToDelete = existingItems.filter(
                (dbItem) => !entity.items.some((item) => item.id === dbItem.id)
            );
    
            // Delete removed items
            for (const item of itemsToDelete) {
                await OrderItemModel.destroy({
                    where: { id: item.id },
                    transaction,
                });
            }
    
            // Upsert (update or insert) current items in entity.items
            for (const item of entity.items) {
                await OrderItemModel.upsert(
                    {
                        id: item.id,
                        name: item.name,
                        price: item.price,
                        product_id: item.productId,
                        quantity: item.quantity,
                        order_id: entity.id,
                    },
                    { transaction }
                );
            }
    
            // Commit the transaction
            await transaction.commit();
        } catch (error) {
            // Rollback the transaction in case of an error
            await transaction.rollback();
            console.error("Error updating order:", error);
            throw new Error("Failed to update order");
        }
    }
    
    async find(id: string): Promise<Order> {
        const orderModel = await OrderModel.findByPk(id, {
            include: [{ model: OrderItemModel, as: 'items' }],
        });
    
        if (!orderModel) {
            throw new Error("Order not found");
        }
    
        return new Order(
            orderModel.id,
            orderModel.customer_id,
            orderModel.items.map((item) => new OrderItem(
                item.id,
                item.name,
                item.price,
                item.product_id,
                item.quantity
            ))
        );
    }

    async findAll(): Promise<Order[]> {
        const orderModels = await OrderModel.findAll({
            include: [{ model: OrderItemModel, as: 'items' }],
        });
        
        //If no orders found should Throw Error
        if (orderModels.length === 0) {
            throw new Error("No orders found");
        }

        return orderModels.map((orderModel) => new Order(
            orderModel.id,
            orderModel.customer_id,
            orderModel.items.map((item) => new OrderItem(
                item.id,
                item.name,
                item.price,
                item.product_id,
                item.quantity
            ))
        ));
    }
}
