import OrderItem from './order_item';

export default class Order {

    private _id: string;
    private _customerId: string;
    private _items: OrderItem[] = [];
    private _total: number = 0;

    constructor(id: string, customerId: string, items: OrderItem[]) {
        this._id = id;
        this._customerId = customerId;
        this._items = items;
        this._total = this.total();
        this.validate();
    }

    validate(): boolean {
        if (this._id.length === 0) {
            throw new Error('Id is required');
        }
        if (this._customerId.length === 0) {
            throw new Error('Customer Id is required');
        }
        if (this._items.length === 0) {
            throw new Error('Order quantity should be greater than 0');
        }
        const itemIds = this._items.map(item => item.id);
        const uniqueItemIds = new Set(itemIds);
        if (itemIds.length !== uniqueItemIds.size) {
            throw new Error("Duplicate OrderItem ID found in the order");
        }
        if (this._items.some(item => item.quantity <= 0)) {
            throw new Error('Items quantity should be greater than 0');
        }

        return true;
    }

    get customerId(): string {
        return this._customerId;
    }

    get items(): OrderItem[] {
        return this._items;
    }

    get id(): string {
        return this._id;
    }

    total(): number {
        return this._items.reduce((acc, item) => acc + item.orderItemTotal(), 0);
    }

    addItem(item: OrderItem): void {
        this._items.push(item);
        this._total = this.total();
    }

    removeItem(itemId: string): void {
        const itemIndex = this._items.findIndex(item => item.id === itemId);
        if (itemIndex === -1) {
            throw new Error('Item not found');
        }
        if (this._items.length === 1) {
            throw new Error('Order quantity should be greater than 0');
        }

        this._items.splice(itemIndex, 1);
        this._total = this.total();
    }

    changeCustomer(customerId: string): void {
        this._customerId = customerId;
    }
}
