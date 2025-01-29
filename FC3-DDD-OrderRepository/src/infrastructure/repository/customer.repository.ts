import Customer from "../../domain/entity/customer";
import Address from "../../domain/entity/address";
import CustomerModel from "../db/sequelize/model/customer.model";
import CustomerRepositoryInterface from "../../domain/repository/customer-repository-interface";

export default class CustomerRepository implements CustomerRepositoryInterface {

    async create(entity: Customer): Promise<void> {
        await CustomerModel.create({
            id: entity.id,
            name: entity.name,
            street: entity.address.street,
            number: entity.address.number,
            zipcode: entity.address.zip,
            city: entity.address.city,
            active: entity.isActive(),
            rewardPoints: entity.rewardPoints,
        });
    }

    async update(entity: Customer): Promise<void> {
        await CustomerModel.update({
            name: entity.name,
            street: entity.address.street,
            number: entity.address.number,
            zipcode: entity.address.zip,
            city: entity.address.city,
            active: entity.isActive(),
            rewardPoints: entity.rewardPoints,
        }, {
            where: {
                id: entity.id,
            },
        });
    }

    async find(id: string): Promise<Customer> {
        const customerModel = await CustomerModel.findByPk(id);
        if (customerModel === null) {
            throw new Error('Customer not found');
        }

        const cust = new Customer(customerModel.id, customerModel.name)
        cust.changeAddress(new Address(
            customerModel.street, 
            customerModel.number, 
            customerModel.zipcode, 
            customerModel.city))
        cust.activate()
        cust.addRewardPoints(customerModel.rewardPoints)
        return cust
    }

    async findAll(): Promise<Customer[]> {
        const customerModels = await CustomerModel.findAll();
        return customerModels.map((customerModel) => {
            const cust = new Customer(customerModel.id, customerModel.name)
            cust.changeAddress(
                new Address(
                    customerModel.street, 
                    customerModel.number, 
                    customerModel.zipcode, 
                    customerModel.city)
                )
            
            cust.activate()
            cust.addRewardPoints(customerModel.rewardPoints)
            return cust
        });
    }

}