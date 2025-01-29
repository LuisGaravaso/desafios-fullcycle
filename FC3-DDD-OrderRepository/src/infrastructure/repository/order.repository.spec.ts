import { Sequelize } from "sequelize-typescript";
import OrderModel from "../db/sequelize/model/order.model";
import CustomerModel from "../db/sequelize/model/customer.model";
import OrderItemModel from "../db/sequelize/model/order_item.model";
import ProductModel from "../db/sequelize/model/product.model";
import CustomerRepository from "./customer.repository";
import Customer from "../../domain/entity/customer";
import Address from "../../domain/entity/address";
import OrderItem from "../../domain/entity/order_item";
import Product from "../../domain/entity/product";
import Order from "../../domain/entity/order";
import OrderRepository from "./order.repository";
import ProductRepository from "./product.repository";

describe("Order Repository test", () => {
    
    let sequelize: Sequelize;

    beforeEach(async () => {
        sequelize = new Sequelize({
            dialect: "sqlite",
            storage: ":memory:",
            logging: false,
            sync: { force: true },
        });

        sequelize.addModels([OrderModel, CustomerModel, OrderItemModel, ProductModel]);
        await sequelize.sync();
    });

    afterEach(async () => {
        await sequelize.close();
    });

    it("should create a new order", async () => {
        const customerRepository = new CustomerRepository();
        const customer = new Customer("Cliente1", "John Doe");

        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer.changeAddress(address);
        await customerRepository.create(customer);

        const productRepository = new ProductRepository();
        const product = new Product("P1", "Product 1", 10);
        await productRepository.create(product);

        const orderItem = new OrderItem(
            "OI1", 
            product.name, 
            product.price, 
            product.id, 
            2);
        
        const order = new Order("O1", customer.id, [orderItem]);
        const orderRepository = new OrderRepository();
        await orderRepository.create(order);

        const orderModel = await OrderModel.findOne(
            { where: { id: "O1" }, 
            include: "items" }
        );

        expect(orderModel).not.toBeNull();
        expect(orderModel.toJSON()).toEqual({
            id: order.id,
            customer_id: customer.id,
            total: order.total(),
            items: [
                {
                    id: orderItem.id,
                    name: orderItem.name,
                    price: orderItem.price,
                    quantity: orderItem.quantity,
                    order_id: order.id,
                    product_id: product.id,
                },
            ],
       });
    });

    it("should update an order by changing the customer", async () => {
        const customerRepository = new CustomerRepository();
        const customer1 = new Customer("Cliente1", "John Doe");

        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer1.changeAddress(address);

        const customer2 = new Customer("Cliente2", "Jane Doe");
        const address2 = new Address("Second Street", 200, "54321", "Springfield");
        customer2.changeAddress(address2);

        await customerRepository.create(customer1);
        await customerRepository.create(customer2);

        const productRepository = new ProductRepository();
        const product = new Product("P1", "Product 1", 10);
        await productRepository.create(product);

        const orderItem = new OrderItem(
            "OI1", 
            product.name, 
            product.price, 
            product.id, 
            2);

        const order = new Order("O1", customer1.id, [orderItem]);
        const orderRepository = new OrderRepository();

        await orderRepository.create(order);

        order.changeCustomer(customer2.id);
        await orderRepository.update(order);

        const orderModel = await OrderModel.findOne(
            { where: { id: "O1" }, 
            include: "items" }
        );

        expect(orderModel).not.toBeNull();
        expect(orderModel.toJSON()).toEqual({
            id: order.id,
            customer_id: customer2.id,
            total: order.total(),
            items: [
                {
                    id: orderItem.id,
                    name: orderItem.name,
                    price: orderItem.price,
                    quantity: orderItem.quantity,
                    order_id: order.id,
                    product_id: product.id,
                },
            ],
        });
    });

    it("should find an order", async () => {
        const customerRepository = new CustomerRepository();
        const customer = new Customer("Cliente1", "John Doe");

        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer.changeAddress(address);
        await customerRepository.create(customer);

        const productRepository = new ProductRepository();
        const product = new Product("P1", "Product 1", 10);
        await productRepository.create(product);

        const orderItem = new OrderItem(
            "OI1", 
            product.name, 
            product.price, 
            product.id, 
            2);

        const order = new Order("O1", customer.id, [orderItem]);
        const orderRepository = new OrderRepository();
        await orderRepository.create(order);

        const foundOrder = await orderRepository.find(order.id);

        expect(foundOrder).toEqual(order);
    });

    it("should throw error when order not found", async () => {
        const orderRepository = new OrderRepository();
        await expect(orderRepository.find("O1")).rejects.toThrow("Order not found");
    });

    it("should find all orders", async () => {
        const customerRepository = new CustomerRepository();
        const customer = new Customer("Cliente1", "John Doe");
        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer.changeAddress(address);
        await customerRepository.create(customer);

        const customer2 = new Customer("Cliente2", "Jane Doe");
        const address2 = new Address("Second Street", 200, "54321", "Springfield");
        customer2.changeAddress(address2);
        await customerRepository.create(customer2);

        const productRepository = new ProductRepository();
        const product = new Product("P1", "Product 1", 10);
        await productRepository.create(product);

        const product2 = new Product("P2", "Product 2", 20);
        await productRepository.create(product2);

        const orderItem = new OrderItem(
            "OI1", 
            product.name, 
            product.price, 
            product.id, 
            2);
        
        const orderItem2 = new OrderItem(
            "OI2", 
            product2.name, 
            product2.price, 
            product2.id, 
            3);
        
        const order = new Order("O1", customer.id, [orderItem]);
        const order2 = new Order("O2", customer2.id, [orderItem2]);
        const orderRepository = new OrderRepository();
        await orderRepository.create(order);
        await orderRepository.create(order2);

        const orders = await orderRepository.findAll();

        expect(orders.length).toBe(2);
        expect(orders[0]).toEqual(order);
        expect(orders[1]).toEqual(order2);
    });

    it("should throw an error when no orders are found", async () => {
        const orderRepository = new OrderRepository();
        await expect(orderRepository.findAll()).rejects.toThrow("No orders found");
    });

    it("should update an order by adding a new OrderItem", async () => {
        const customerRepository = new CustomerRepository();
        const customer = new Customer("Cliente1", "John Doe");

        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer.changeAddress(address);
        await customerRepository.create(customer);

        const productRepository = new ProductRepository();
        const product = new Product("P1", "Product 1", 10);
        await productRepository.create(product);

        const orderItem = new OrderItem(
            "OI1", 
            product.name, 
            product.price, 
            product.id, 
            2);

        const order = new Order("O1", customer.id, [orderItem]);
        const orderRepository = new OrderRepository();
        await orderRepository.create(order);

        const newProduct = new Product("P2", "Product 2", 20);
        await productRepository.create(newProduct);
        
        const newOrderItem = new OrderItem(
            "OI2", 
            newProduct.name, 
            newProduct.price, 
            newProduct.id, 
            3);

        order.addItem(newOrderItem);
        await orderRepository.update(order);
        
        const orderModel = await OrderModel.findOne(
            { where: { id: "O1" }, 
            include: "items" }
        );

        expect(orderModel).not.toBeNull();
        expect(orderModel.toJSON()).toEqual({
            id: order.id,
            customer_id: customer.id,
            total: order.total(),
            items: [
                {
                    id: orderItem.id,
                    name: orderItem.name,
                    price: orderItem.price,
                    quantity: orderItem.quantity,
                    order_id: order.id,
                    product_id: product.id,
                },
                {
                    id: newOrderItem.id,
                    name: newOrderItem.name,
                    price: newOrderItem.price,
                    quantity: newOrderItem.quantity,
                    order_id: order.id,
                    product_id: newProduct.id,
                },
            ],
        });
    });

    it("should update an order by removing an OrderItem", 
        async () => {
        const customerRepository = new CustomerRepository();
        const customer = new Customer("Cliente1", "John Doe");

        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer.changeAddress(address);
        await customerRepository.create(customer);

        const productRepository = new ProductRepository();
        const product1 = new Product("P1", "Product 1", 10);
        await productRepository.create(product1);
        
        const product2 = new Product("P2", "Product 2", 20);
        await productRepository.create(product2);

        const orderItem1 = new OrderItem(
            "OI1", 
            product1.name, 
            product1.price, 
            product1.id, 
            2);
        
        const orderItem2 = new OrderItem(
            "OI2", 
            product2.name, 
            product2.price, 
            product2.id, 
            3);

        const order = new Order("O1", customer.id, [orderItem1, orderItem2]);
        const orderRepository = new OrderRepository();
        await orderRepository.create(order);
        
        order.removeItem(orderItem1.id);
        await orderRepository.update(order);

        const orderModel = await OrderModel.findOne(
            { where: { id: "O1" }, 
            include: "items" }
        );

        expect(orderModel).not.toBeNull();

        expect(orderModel.toJSON()).toEqual({
            id: order.id,
            customer_id: customer.id,
            total: order.total(),
            items: [
                {
                    id: orderItem2.id,
                    name: orderItem2.name,
                    price: orderItem2.price,
                    quantity: orderItem2.quantity,
                    order_id: order.id,
                    product_id: product2.id,
                },
            ],
        });
    });

});