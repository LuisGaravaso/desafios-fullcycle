import Customer from "../entity/customer";
import Order from "../entity/order";
import OrderItem from "../entity/order_item";
import OrderService from "./order.service";

describe("Order Service Unit Test", () => {
    
    it("should get total of all orders", () => {
        
        const item1 = new OrderItem("1", "item1", 10, "p1", 2);
        const item2 = new OrderItem("2", "item2", 20, "p2", 3);
        const item3 = new OrderItem("3", "item3", 30, "p3", 4);
        const item4 = new OrderItem("4", "item4", 40, "p4", 5);

        const order1 = new Order("1", "1", [item1, item2]);
        const order2 = new Order("2", "2", [item3, item4]);
        const orders = [order1, order2];

        expect(OrderService.total(orders)).toBe(400);

    });

    it("should place an order", () => {
        
        const customer = new Customer("1", "customer1");
        const item1 = new OrderItem("1", "item1", 10, "p1", 2);
        const item2 = new OrderItem("2", "item2", 20, "p2", 3);
        const items = [item1, item2];

        const orderPlaced = OrderService.placeOrder(customer, items);
        expect(orderPlaced.total()).toBe(80);
        expect(customer.rewardPoints).toBe(40);
        



    }); 
});