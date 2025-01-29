import Customer from "./customer";
import Address from "./address";

describe("Customer unit tests", () => {
  
    it("should throw error when id is empty", () => {
        expect(() => new Customer("", "John")).toThrow('Id is required');
    });

    it("should throw error when name is empty", () => {
        expect(() => new Customer("1", "")).toThrow('Name is required');
    });

    it("should should change name", () => {
        const customer = new Customer("1", "John");
        customer.changeName("Doe");
        expect(customer.name).toBe("Doe");
    });

    it("should activate customer", () => {
        const customer = new Customer("1", "John");
        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer.changeAddress(address);
        customer.activate();
        expect(customer.isActive()).toBe(true);
    });

    it("should deactivate customer", () => {
        const customer = new Customer("1", "John");
        customer.deactivate();
        expect(customer.isActive()).toBe(false);
    });

    it("should throw error when activating customer without address", () => {
        const customer = new Customer("1", "John");
        expect(() => customer.activate()).toThrow('Address is mandatory to activate customer');
    });

    it("should add reward points", () => {
        const customer = new Customer("1", "John");
        expect(customer.rewardPoints).toBe(0);
        customer.addRewardPoints(100);
        expect(customer.rewardPoints).toBe(100);
        customer.addRewardPoints(100);
        expect(customer.rewardPoints).toBe(200);
    });
});
