import { ChatService } from '../../services/chat/chat.service';
import { Test, TestingModule } from '@nestjs/testing';
import { ChatController } from './chat.controller';

describe('ChatController', () => {
  let controller: ChatController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [ChatController],
      providers:[ChatService]
    }).compile();

    controller = module.get<ChatController>(ChatController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
