package io.cecg.referenceapplication.domain.repository.dto;

import lombok.*;
import lombok.extern.slf4j.Slf4j;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Slf4j
@EqualsAndHashCode(of = "name")
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "counter")
@Builder
@Getter
@Setter
public class Counter {
    @Id
    private String name;

    @Column(nullable = false)
    private Long counter;

    public void incrementCounter() {
        this.counter++;
    }
}
