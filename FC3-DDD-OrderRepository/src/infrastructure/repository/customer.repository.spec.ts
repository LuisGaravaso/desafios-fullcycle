import { Sequelize } from "sequelize-typescript";
import Customer from "../../domain/entity/customer";
import Address from "../../domain/entity/address";
import CustomerModel from "../db/sequelize/model/customer.model";
import CustomerRepository from "./customer.repository";

describe("Customer Repository test", () => {
    
    let sequelize: Sequelize;

    beforeEach(async () => {
        sequelize = new Sequelize({
            dialect: "sqlite",
            storage: ":memory:",
            logging: false,
            sync: { force: true },
        });

        sequelize.addModels([CustomerModel]);
        await sequelize.sync();
    });

    afterEach(async () => {
        await sequelize.close();
    });

    it("should create a customer", async () => {

        const customerRepository = new CustomerRepository();
        const customer = new Customer("1", "John Doe");
        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer.changeAddress(address);
        await customerRepository.create(customer);

        const customerModel = await CustomerModel.findOne({ where: { id: "1" } });

        expect(customerModel).not.toBeNull();
        expect(customerModel.toJSON()).toEqual({
            id: customer.id,
            name: customer.name,
            street: customer.address.street,
            number: customer.address.number,
            zipcode: customer.address.zip,
            city: customer.address.city,
            active: true,
            rewardPoints: 0,
        });
    });

    it("should update a customer", async () => {

        const customerRepository = new CustomerRepository();
        const customer = new Customer("1", "John Doe");
        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer.changeAddress(address);
        await customerRepository.create(customer);

        customer.changeName("Jane Doe");
        customer.deactivate();
        customer.addRewardPoints(100);
        await customerRepository.update(customer);

        const customerModel = await CustomerModel.findOne({ where: { id: "1" } });

        expect(customerModel).not.toBeNull();
        expect(customerModel.toJSON()).toEqual({
            id: customer.id,
            name: customer.name,
            street: customer.address.street,
            number: customer.address.number,
            zipcode: customer.address.zip,
            city: customer.address.city,
            active: false,
            rewardPoints: 100,
        });
    });

    it("should find a customer", async () => {

        const customerRepository = new CustomerRepository();
        const customer = new Customer("1", "John Doe");
        const address = new Address("Main Street", 100, "12345", "Springfield");
        customer.changeAddress(address);
        await customerRepository.create(customer);

        const foundCustomer = await customerRepository.find("1");

        expect(foundCustomer.id).toBe(customer.id);
        expect(foundCustomer.name).toBe(customer.name);
        expect(foundCustomer.address.street).toBe(customer.address.street);
        expect(foundCustomer.address.number).toBe(customer.address.number);
        expect(foundCustomer.address.zip).toBe(customer.address.zip);
        expect(foundCustomer.address.city).toBe(customer.address.city);
        expect(foundCustomer.isActive()).toBe(true);
        expect(foundCustomer.rewardPoints).toBe(0);
    });

    it("should throw error when customer not found", async () => {

        const customerRepository = new CustomerRepository();
        await expect(customerRepository.find("1")).rejects.toThrow("Customer not found");
    });

    it("should find all customers", async () => {

        const customerRepository = new CustomerRepository();
        const customer1 = new Customer("1", "John Doe");
        const address1 = new Address("Main Street", 100, "12345", "Springfield");
        customer1.changeAddress(address1);
        await customerRepository.create(customer1);

        const customer2 = new Customer("2", "Jane Doe");
        const address2 = new Address("Second Street", 200, "54321", "Springfield");
        customer2.changeAddress(address2);
        await customerRepository.create(customer2);

        const customers = await customerRepository.findAll();

        expect(customers.length).toBe(2);

        const foundCustomer1 = customers.find((cust) => cust.id === customer1.id);
        expect(foundCustomer1).not.toBeUndefined();
        expect(foundCustomer1!.id).toBe(customer1.id);
        expect(foundCustomer1!.name).toBe(customer1.name);
        expect(foundCustomer1!.address.street).toBe(customer1.address.street);
        expect(foundCustomer1!.address.number).toBe(customer1.address.number);
        expect(foundCustomer1!.address.zip).toBe(customer1.address.zip);
        expect(foundCustomer1!.address.city).toBe(customer1.address.city);
        expect(foundCustomer1!.isActive()).toBe(true);
        expect(foundCustomer1!.rewardPoints).toBe(0);

        const foundCustomer2 = customers.find((cust) => cust.id === customer2.id);
        expect(foundCustomer2).not.toBeUndefined();
        expect(foundCustomer2!.id).toBe(customer2.id);
        expect(foundCustomer2!.name).toBe(customer2.name);
        expect(foundCustomer2!.address.street).toBe(customer2.address.street);
        expect(foundCustomer2!.address.number).toBe(customer2.address.number);
        expect(foundCustomer2!.address.zip).toBe(customer2.address.zip);
        expect(foundCustomer2!.address.city).toBe(customer2.address.city);
        expect(foundCustomer2!.isActive()).toBe(true);
        expect(foundCustomer2!.rewardPoints).toBe(0);
    });

});