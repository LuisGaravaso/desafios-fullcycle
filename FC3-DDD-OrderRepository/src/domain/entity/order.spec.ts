import Order from "./order";
import OrderItem from "./order_item";

describe("Order unit tests", () => {
  
    it("should throw error when id is empty", () => {
        expect(() => new Order("", "1", [])).toThrow('Id is required');
    });

    it("should throw error when customer id is empty", () => {
        expect(() => new Order("1", "", [])).toThrow('Customer Id is required');
    });

    it("should throw error when order items is empty", () => {
        expect(() => new Order("1", "1", [])).toThrow('Order quantity should be greater than 0');
    });

    it("should calculate total", () => {
        const item1 = new OrderItem("1", "item1", 10, "p1", 2);
        const item2 = new OrderItem("2", "item2", 20, "p2", 3);
        const items = [item1, item2];
        const order = new Order("1", "1", items);
        expect(order.total()).toBe(80);
    });

    it("should throw error if item qte is less or equal to 0", () => {
        const item1 = new OrderItem("1", "item1", 10, "p1", 0);
        const items = [item1];
        expect(() => new Order("1", "1", items)).toThrow('Items quantity should be greater than 0');
    });
    
    it("should throw error if item id is duplicated", () => {
        const item1 = new OrderItem("1", "item1", 10, "p1", 2);
        const item2 = new OrderItem("1", "item2", 20, "p2", 3);
        const items = [item1, item2];
        expect(() => new Order("1", "1", items)).toThrow('Duplicate OrderItem ID found in the order');
    }); 
});
