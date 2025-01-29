import Product from "./product";

describe("Product unit tests", () => {

    it("should throw error when id is empty", () => {
        expect(() => {
            const product = new Product("", "item1", 10);

        }).toThrow('Id is required');
    });

    it("should throw error when name is empty", () => {
        expect(() => {
            const product = new Product("1", "", 10);

        }).toThrow('Name is required');
    });

    it("should throw error when price is 0", () => {
        expect(() => {
            const product = new Product("1", "item1", 0);

        }).toThrow('Price should be greater than 0');
    });

    it("should change name", () => {
        const product = new Product("1", "item1", 10);
        product.changeName("item2");
        expect(product.name).toBe("item2");
    });
   
    it("should change price", () => {
        const product = new Product("1", "item1", 10);
        product.changePrice(20);
        expect(product.price).toBe(20);
    });
});
