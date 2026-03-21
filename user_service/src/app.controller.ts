import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';
import axios from 'axios';
@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  @Get('/api/user')
  getUser(): any {
    return [
      {
        id: 1,
        name: 'John Doe',
        email: 'jd@example.com'
      }
    ];
  }

  @Get('/api/user/payment')
  async getUserPayment(): Promise<any> {
    try {
      const response = await axios.get('http://payment_service:8080/api/payment/user');
      return response.data;
    } catch (error) {
      return { error: 'Failed to fetch payment info' };
    }
  }
}
